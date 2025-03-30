package godot

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/bashidogames/gevm/config"
	"github.com/bashidogames/gevm/internal/archiving"
	"github.com/bashidogames/gevm/internal/downloading"
	"github.com/bashidogames/gevm/internal/environment"
	"github.com/bashidogames/gevm/internal/locator"
	"github.com/bashidogames/gevm/internal/utils"
	"github.com/bashidogames/gevm/semver"
	"github.com/jedib0t/go-pretty/v6/table"
)

const CACHE_FOLDER = "godot"

type Service struct {
	Environment *environment.Environment
	Locator     *locator.Locator
	Config      *config.Config
}

func (s *Service) Download(semver semver.Semver) error {
	s.Config.Logger.Debug("Attempting to download '%s' godot...", semver.GodotString())

	asset, err := s.Environment.FetchGodotAsset(semver)
	if errors.Is(err, downloading.ErrNotFound) {
		s.Config.Logger.Error("Godot '%s' not found. Use 'gevm versions list' to see available versions.", semver.GodotString())
		return nil
	}
	if err != nil {
		return fmt.Errorf("fetch asset failed: %w", err)
	}

	archivePath := s.archivePath(asset.Name)

	exists, err := utils.DoesExist(archivePath)
	if err != nil {
		return fmt.Errorf("failed to check existence: %w", err)
	}

	if exists {
		s.Config.Logger.Info("Godot '%s' already downloaded", semver.GodotString())
		return nil
	}

	s.Config.Logger.Debug("Downloading from: %s", asset.DownloadURL)
	s.Config.Logger.Debug("Downloading to: %s", archivePath)

	err = downloading.Download(s.Config.Logger, asset.DownloadURL, archivePath, s.Config.Silent)
	if errors.Is(err, downloading.ErrNotFound) {
		s.Config.Logger.Error("Godot '%s' not found. Use 'gevm versions list' to see available versions.", semver.GodotString())
		return nil
	}
	if err != nil {
		return fmt.Errorf("download failed: %w", err)
	}

	s.Config.Logger.Info("Godot '%s' downloaded", semver.GodotString())
	return nil
}

func (s *Service) Uninstall(semver semver.Semver, logMissing bool) error {
	s.Config.Logger.Debug("Attempting to uninstall '%s' godot...", semver.GodotString())

	targetDirectory := s.targetDirectory(semver)

	exists, err := utils.DoesExist(targetDirectory)
	if err != nil {
		return fmt.Errorf("failed to check existence: %w", err)
	}

	if !exists {
		if logMissing {
			s.Config.Logger.Error("Godot '%s' not found", semver.GodotString())
		}

		return nil
	}

	s.Config.Logger.Debug("Removing directory: %s", targetDirectory)

	err = os.RemoveAll(targetDirectory)
	if err != nil {
		return fmt.Errorf("cannot remove target directory: %w", err)
	}

	s.Config.Logger.Info("Godot '%s' uninstalled", semver.GodotString())
	return nil
}

func (s *Service) Install(semver semver.Semver) error {
	s.Config.Logger.Debug("Attempting to install '%s' godot...", semver.GodotString())

	asset, err := s.Environment.FetchGodotAsset(semver)
	if errors.Is(err, downloading.ErrNotFound) {
		s.Config.Logger.Error("Godot '%s' not found. Use 'gevm versions list' to see available versions.", semver.GodotString())
		return nil
	}
	if err != nil {
		return fmt.Errorf("fetch asset failed: %w", err)
	}

	targetDirectory := s.targetDirectory(semver)
	archivePath := s.archivePath(asset.Name)

	exists, err := utils.DoesExist(targetDirectory)
	if err != nil {
		return fmt.Errorf("failed to check existence: %w", err)
	}

	if exists {
		s.Config.Logger.Info("Godot '%s' already installed", semver.GodotString())
		return nil
	}

	err = os.MkdirAll(s.Config.GodotRootDirectory, utils.OS_DIRECTORY)
	if err != nil {
		return fmt.Errorf("cannot make directory: %w", err)
	}

	err = os.RemoveAll(targetDirectory)
	if err != nil {
		return fmt.Errorf("cannot remove target directory: %w", err)
	}

	s.Config.Logger.Debug("Downloading from: %s", asset.DownloadURL)
	s.Config.Logger.Debug("Downloading to: %s", archivePath)

	err = downloading.Download(s.Config.Logger, asset.DownloadURL, archivePath, s.Config.Silent)
	if errors.Is(err, downloading.ErrNotFound) {
		s.Config.Logger.Error("Godot '%s' not found. Use 'gevm versions list' to see available versions.", semver.GodotString())
		return nil
	}
	if err != nil {
		return fmt.Errorf("download failed: %w", err)
	}

	s.Config.Logger.Debug("Unzipping from: %s", archivePath)
	s.Config.Logger.Debug("Unzipping to: %s", targetDirectory)

	err = archiving.Unzip(s.Config.Logger, archivePath, targetDirectory)
	if err != nil {
		return fmt.Errorf("unzip failed: %w", err)
	}

	s.Config.Logger.Info("Godot '%s' installed", semver.GodotString())
	return nil
}

func (s *Service) Use(semver semver.Semver) error {
	s.Config.Logger.Debug("Attempting to use '%s' godot...", semver.GodotString())

	targetPath, err := s.Locator.TargetPath(semver)
	if errors.Is(err, os.ErrNotExist) {
		s.Config.Logger.Error("Godot '%s' not found. Use `gevm godot list` to see installed versions.", semver.GodotString())
		return nil
	}
	if err != nil {
		return fmt.Errorf("cannot determine target path: %w", err)
	}

	linkPath := s.Locator.LinkPath(semver)

	s.Config.Logger.Debug("Creating godot symlink: %s => %s", linkPath, targetPath)

	err = os.MkdirAll(s.Config.BinDirectory, utils.OS_DIRECTORY)
	if err != nil {
		return fmt.Errorf("cannot make directory: %w", err)
	}

	err = os.RemoveAll(linkPath)
	if err != nil {
		return fmt.Errorf("cannot remove link path: %w", err)
	}

	err = os.Symlink(targetPath, linkPath)
	if err != nil {
		return fmt.Errorf("cannot create link: %w", err)
	}

	s.Config.Logger.Info("Using '%s' godot", semver.GodotString())
	return nil
}

func (s *Service) Path(semver semver.Semver) error {
	targetPath, err := s.Locator.TargetPath(semver)
	if errors.Is(err, os.ErrNotExist) {
		s.Config.Logger.Error("Godot '%s' not found. Use `gevm godot list` to see installed versions.", semver.GodotString())
		return nil
	}
	if err != nil {
		return fmt.Errorf("cannot determine target path: %w", err)
	}

	utils.Printlnf(targetPath)
	return nil
}

func (s *Service) List() error {
	entries, err := os.ReadDir(s.Config.GodotRootDirectory)
	if !errors.Is(err, os.ErrNotExist) && err != nil {
		return fmt.Errorf("cannot read godot root directory: %w", err)
	}

	if len(entries) == 0 {
		s.Config.Logger.Info("No godot engine versions installed")
		return nil
	}

	t := table.NewWriter()
	t.AppendHeader(table.Row{"Version", "Release", "Mono?"})

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		semver, err := semver.Parse(entry.Name())
		if err != nil {
			s.Config.Logger.Warning("Failed to recognize version: %s", err)
			continue
		}

		version := semver.Relver.Version.String()
		release := semver.Relver.Release.String()
		mono := semver.Mono

		t.AppendRow(table.Row{version, release, mono})
	}

	t.SetOutputMirror(os.Stdout)
	t.Render()

	return nil
}

func (s *Service) targetDirectory(semver semver.Semver) string {
	return filepath.Join(s.Config.GodotRootDirectory, semver.GodotString())
}

func (s *Service) archivePath(name string) string {
	return filepath.Join(s.Config.CacheDirectory, CACHE_FOLDER, name)
}

func New(environment *environment.Environment, locator *locator.Locator, config *config.Config) *Service {
	return &Service{
		Environment: environment,
		Locator:     locator,
		Config:      config,
	}
}
