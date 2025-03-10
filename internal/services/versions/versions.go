package versions

import (
	"fmt"
	"os"

	"github.com/bashidogames/gevm/config"
	"github.com/bashidogames/gevm/internal/environment"
	"github.com/bashidogames/gevm/internal/platform"
	"github.com/jedib0t/go-pretty/v6/table"
)

type Service struct {
	Environment *environment.Environment
	Config      *config.Config
}

func (s *Service) Detailed(all bool, mono bool) error {
	downloads, err := s.Environment.FetchDownloads(mono)
	if err != nil {
		return fmt.Errorf("cannot fetch environment downloads: %w", err)
	}

	header := table.Row{"Version", "Release"}
	for _, platform := range platform.Platforms {
		header = append(header, platform)
	}

	t := table.NewWriter()
	t.AppendHeader(header)

	for _, download := range downloads {
		if !download.Relver.IsStable() && !all {
			continue
		}

		availabilities := map[platform.Platform]bool{}
		for _, platform := range platform.Platforms {
			availabilities[platform] = download.HasAsset(platform)
		}

		row := table.Row{
			download.Relver.Version,
			download.Relver.Release,
		}

		for _, platform := range platform.Platforms {
			row = append(row, availabilities[platform])
		}

		t.AppendRow(row)
	}

	t.SetOutputMirror(os.Stdout)
	t.Render()

	return nil
}

func (s *Service) List(all bool, mono bool) error {
	downloads, err := s.Environment.FetchDownloads(mono)
	if err != nil {
		return fmt.Errorf("cannot fetch environment downloads: %w", err)
	}

	t := table.NewWriter()
	t.AppendHeader(table.Row{"Version", "Release", s.Config.Platform})

	for _, download := range downloads {
		if !download.Relver.IsStable() && !all {
			continue
		}

		available := download.HasAsset(s.Config.Platform)
		row := table.Row{
			download.Relver.Version,
			download.Relver.Release,
			available,
		}

		t.AppendRow(row)
	}

	t.SetOutputMirror(os.Stdout)
	t.Render()

	return nil
}

func New(environment *environment.Environment, config *config.Config) *Service {
	return &Service{
		Environment: environment,
		Config:      config,
	}
}
