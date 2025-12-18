package gevm

import (
	"fmt"

	"github.com/bashidogames/gevm/config"
	"github.com/bashidogames/gevm/internal/environment"
	"github.com/bashidogames/gevm/internal/environment/fetcher"
	"github.com/bashidogames/gevm/internal/environment/github"
	"github.com/bashidogames/gevm/internal/locator"
	"github.com/bashidogames/gevm/internal/services/cache"
	"github.com/bashidogames/gevm/internal/services/exporttemplates"
	"github.com/bashidogames/gevm/internal/services/godot"
	"github.com/bashidogames/gevm/internal/services/settings"
	"github.com/bashidogames/gevm/internal/services/versions"
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
