package desktop

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
	s.Config.Logger.Debug("Attempting to remove '%s' desktop shortcut...", semver.GodotString())

	shortcutPath := s.Locator.DesktopShortcutPath(semver)

	s.Config.Logger.Debug("Removing desktop shortcut: %s", shortcutPath)

	exists, err := utils.DoesExist(shortcutPath)
	if err != nil {
		return fmt.Errorf("failed to check existence: %w", err)
	}

	if !exists {
		if logMissing {
			s.Config.Logger.Error("Desktop shortcut '%s' not found", semver.GodotString())
		}

		return nil
	}

	err = os.RemoveAll(shortcutPath)
	if err != nil {
		return fmt.Errorf("cannot remove shortcut: %w", err)
	}

	s.Config.Logger.Info("Desktop '%s' shortcut removed", semver.GodotString())
	return nil
}

func (s *Service) Add(semver semver.Semver) error {
	s.Config.Logger.Debug("Attempting to add '%s' desktop shortcut...", semver.GodotString())

	targetPath, err := s.Locator.TargetPath(semver)
	if errors.Is(err, os.ErrNotExist) {
		s.Config.Logger.Error("Godot '%s' not found. Use `gevm godot list` to see installed versions.", semver.GodotString())
		return nil
	}
	if err != nil {
		return fmt.Errorf("cannot determine target path: %w", err)
	}

	shortcutPath := s.Locator.DesktopShortcutPath(semver)
	shortcutName := s.Locator.ShortcutName(semver)

	s.Config.Logger.Debug("Adding '%s' desktop shortcut: %s => %s", shortcutName, shortcutPath, targetPath)

	exists, err := utils.DoesExist(shortcutPath)
	if err != nil {
		return fmt.Errorf("failed to check existence: %w", err)
	}

	if exists {
		s.Config.Logger.Info("Desktop shortcut '%s' already added", semver.GodotString())
		return nil
	}

	err = os.MkdirAll(s.Config.DesktopShortcutDirectory, utils.OS_DIRECTORY)
	if err != nil {
		return fmt.Errorf("cannot make directory: %w", err)
	}

	err = shortcut.Create(shortcutPath, targetPath, shortcutName)
	if err != nil {
		return fmt.Errorf("cannot create shortcut: %w", err)
	}

	s.Config.Logger.Info("Desktop '%s' shortcut added", semver.GodotString())
	return nil
}

func New(locator *locator.Locator, config *config.Config) *Service {
	return &Service{
		Locator: locator,
		Config:  config,
	}
}
