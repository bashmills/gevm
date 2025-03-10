package application

import (
	"errors"
	"fmt"
	"os"

	"github.com/bashidogames/gevm/config"
	"github.com/bashidogames/gevm/internal/locator"
	"github.com/bashidogames/gevm/internal/shortcut"
	"github.com/bashidogames/gevm/internal/utils"
	"github.com/bashidogames/gevm/semver"
)

type Service struct {
	Locator *locator.Locator
	Config  *config.Config
}

func (s *Service) Remove(semver semver.Semver, logMissing bool) error {
	s.Config.Logger.Trace("Attempting to remove '%s' application shortcut...", semver.GodotString())

	shortcutPath := s.Locator.ApplicationShortcutPath(semver)

	s.Config.Logger.Trace("Removing application shortcut: %s", shortcutPath)

	exists, err := utils.DoesExist(shortcutPath)
	if err != nil {
		return fmt.Errorf("failed to check existence: %w", err)
	}

	if !exists {
		if logMissing {
			s.Config.Logger.Error("Application shortcut '%s' not found", semver.GodotString())
		}

		return nil
	}

	err = os.RemoveAll(shortcutPath)
	if err != nil {
		return fmt.Errorf("cannot remove shortcut: %w", err)
	}

	s.Config.Logger.Info("Application '%s' shortcut removed", semver.GodotString())
	return nil
}

func (s *Service) Add(semver semver.Semver) error {
	s.Config.Logger.Trace("Attempting to add '%s' application shortcut...", semver.GodotString())

	targetPath, err := s.Locator.TargetPath(semver)
	if errors.Is(err, os.ErrNotExist) {
		s.Config.Logger.Error("Godot '%s' not found. Use `gevm godot list` to see installed versions.", semver.GodotString())
		return nil
	}
	if err != nil {
		return fmt.Errorf("cannot determine target path: %w", err)
	}

	shortcutPath := s.Locator.ApplicationShortcutPath(semver)
	shortcutName := s.Locator.ShortcutName(semver)

	s.Config.Logger.Trace("Adding '%s' application shortcut: %s => %s", shortcutName, shortcutPath, targetPath)

	exists, err := utils.DoesExist(shortcutPath)
	if err != nil {
		return fmt.Errorf("failed to check existence: %w", err)
	}

	if exists {
		s.Config.Logger.Info("Application shortcut '%s' already added", semver.GodotString())
		return nil
	}

	err = shortcut.Create(shortcutPath, targetPath, shortcutName)
	if err != nil {
		return fmt.Errorf("cannot create shortcut: %w", err)
	}

	s.Config.Logger.Info("Application '%s' shortcut added", semver.GodotString())
	return nil
}

func New(locator *locator.Locator, config *config.Config) *Service {
	return &Service{
		Locator: locator,
		Config:  config,
	}
}
