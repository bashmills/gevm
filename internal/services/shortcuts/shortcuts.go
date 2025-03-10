package shortcuts

import (
	"github.com/bashidogames/gevm/config"
	"github.com/bashidogames/gevm/internal/locator"
	"github.com/bashidogames/gevm/internal/services/shortcuts/application"
	"github.com/bashidogames/gevm/internal/services/shortcuts/desktop"
)

type Service struct {
	Application *application.Service
	Desktop     *desktop.Service
}

func New(locator *locator.Locator, config *config.Config) *Service {
	return &Service{
		Application: application.New(locator, config),
		Desktop:     desktop.New(locator, config),
	}
}
