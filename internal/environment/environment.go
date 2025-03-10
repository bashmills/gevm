package environment

import (
	"fmt"
	"slices"

	"github.com/bashidogames/gevm/config"
	"github.com/bashidogames/gevm/internal/environment/fetcher"
	"github.com/bashidogames/gevm/internal/platform"
	"github.com/bashidogames/gevm/internal/repository"
	"github.com/bashidogames/gevm/semver"
)

type Environment struct {
	Fetcher fetcher.Fetcher
	Config  *config.Config
}

func (e *Environment) FetchExportTemplatesAsset(semver semver.Semver) (*repository.Asset, error) {
	asset, err := e.Fetcher.FetchAsset(platform.ExportTemplates, semver)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch export templates asset: %w", err)
	}

	return asset, nil
}

func (e *Environment) FetchGodotAsset(semver semver.Semver) (*repository.Asset, error) {
	asset, err := e.Fetcher.FetchAsset(e.Config.Platform, semver)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch godot asset: %w", err)
	}

	return asset, nil
}

func (e *Environment) FetchDownloads(mono bool) ([]repository.Download, error) {
	downloads, err := e.Fetcher.FetchDownloads(mono)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch downloads: %w", err)
	}

	slices.SortFunc(downloads, func(a repository.Download, b repository.Download) int { return a.Relver.Compare(b.Relver) })

	return downloads, nil
}

func New(fetcher fetcher.Fetcher, config *config.Config) (*Environment, error) {
	return &Environment{
		Fetcher: fetcher,
		Config:  config,
	}, nil
}
