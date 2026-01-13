package gevm

import (
	"fmt"

	"github.com/bashmills/gevm/config"
	"github.com/bashmills/gevm/internal/environment"
	"github.com/bashmills/gevm/internal/environment/fetcher"
	"github.com/bashmills/gevm/internal/environment/github"
	"github.com/bashmills/gevm/internal/locator"
	"github.com/bashmills/gevm/internal/services/cache"
	"github.com/bashmills/gevm/internal/services/exporttemplates"
	"github.com/bashmills/gevm/internal/services/godot"
	"github.com/bashmills/gevm/internal/services/settings"
	"github.com/bashmills/gevm/internal/services/versions"
)

type App struct {
	Versions        *versions.Service
	ExportTemplates *exporttemplates.Service
	Godot           *godot.Service
	Settings        *settings.Service
	Cache           *cache.Service
}

func New(config *config.Config) (*App, error) {
	environment, err := environment.New([]fetcher.Fetcher{github.New(config)}, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create environment: %w", err)
	}

	locator, err := locator.New(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create locator: %w", err)
	}

	versionsService := versions.New(environment, config)
	exportTemplatesService := exporttemplates.New(environment, config)
	godotService := godot.New(environment, exportTemplatesService, locator, config)
	settingsService := settings.New(config)
	cacheService := cache.New(config)

	return &App{
		Versions:        versionsService,
		ExportTemplates: exportTemplatesService,
		Godot:           godotService,
		Settings:        settingsService,
		Cache:           cacheService,
	}, nil
}
