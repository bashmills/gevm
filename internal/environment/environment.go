package environment

import (
	"errors"
	"fmt"
	"slices"

	"github.com/bashmills/gevm/config"
	"github.com/bashmills/gevm/internal/downloading"
	"github.com/bashmills/gevm/internal/environment/fetcher"
	"github.com/bashmills/gevm/internal/platform"
	"github.com/bashmills/gevm/internal/repository"
	"github.com/bashmills/gevm/semver"
)

type Environment struct {
	Fetchers []fetcher.Fetcher
	Config   *config.Config
}

func (e *Environment) FetchExportTemplatesAsset(semver semver.Semver) (*repository.Asset, error) {
	var result *repository.Asset
	for _, fetcher := range e.Fetchers {
		asset, err := fetcher.FetchAsset(platform.ExportTemplates, semver)
		if errors.Is(err, downloading.ErrNotFound) {
			continue
		}
		if err != nil {
			return nil, fmt.Errorf("failed to fetch export templates asset: %w", err)
		}

		result = asset
		break
	}

	if result == nil {
		return nil, downloading.ErrNotFound
	}

	return result, nil
}

func (e *Environment) FetchGodotAsset(semver semver.Semver) (*repository.Asset, error) {
	var result *repository.Asset
	for _, fetcher := range e.Fetchers {
		asset, err := fetcher.FetchAsset(e.Config.Platform, semver)
		if errors.Is(err, downloading.ErrNotFound) {
			continue
		}
		if err != nil {
			return nil, fmt.Errorf("failed to fetch godot asset: %w", err)
		}

		result = asset
		break
	}

	if result == nil {
		return nil, downloading.ErrNotFound
	}

	return result, nil
}

func (e *Environment) FetchDownloads(mono bool) ([]repository.Download, error) {
	var result []repository.Download
	for _, fetcher := range e.Fetchers {
		downloads, err := fetcher.FetchDownloads(mono)
		if errors.Is(err, downloading.ErrNotFound) {
			continue
		}
		if err != nil {
			return nil, fmt.Errorf("failed to fetch downloads: %w", err)
		}

		result = downloads
		break
	}

	if result == nil {
		return nil, downloading.ErrNotFound
	}

	slices.SortFunc(result, func(a repository.Download, b repository.Download) int { return a.Relver.Compare(b.Relver) })

	return result, nil
}

func New(fetchers []fetcher.Fetcher, config *config.Config) (*Environment, error) {
	return &Environment{
		Fetchers: fetchers,
		Config:   config,
	}, nil
}
