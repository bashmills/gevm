package godot

import (
	"fmt"

	"github.com/bashmills/gevm"
	"github.com/bashmills/gevm/semver"
)

type Download struct {
	Version                string `arg:"" help:"Godot engine version to download to cache in the format x.x.x.x, x.x.x or x.x"`
	ExcludeExportTemplates bool   `short:"e" help:"Exclude export templates in download"`
	Release                string `short:"r" default:"stable" help:"Release to use (dev1, alpha2, beta3, rc4, stable, etc)"`
	Mono                   bool   `short:"m" help:"Use mono version"`
}

func (c *Download) Run(app *gevm.App) error {
	if !c.ExcludeExportTemplates {
		err := app.ExportTemplates.Download(semver.Maybe(c.Version, c.Release, c.Mono))
		if err != nil {
			return fmt.Errorf("cannot download export templates: %w", err)
		}
	}

	err := app.Godot.Download(semver.Maybe(c.Version, c.Release, c.Mono))
	if err != nil {
		return fmt.Errorf("cannot download godot: %w", err)
	}

	return nil
}

type Uninstall struct {
	Version                string `arg:"" help:"Godot engine version to uninstall in the format x.x.x.x, x.x.x or x.x"`
	ExcludeExportTemplates bool   `short:"e" help:"Exclude export templates in uninstall"`
	Release                string `short:"r" default:"stable" help:"Release to use (dev1, alpha2, beta3, rc4, stable, etc)"`
	Mono                   bool   `short:"m" help:"Use mono version"`
}

func (c *Uninstall) Run(app *gevm.App) error {
	if !c.ExcludeExportTemplates {
		err := app.ExportTemplates.Uninstall(semver.Maybe(c.Version, c.Release, c.Mono), false)
		if err != nil {
			return fmt.Errorf("cannot uninstall export templates: %w", err)
		}
	}

	err := app.Godot.Uninstall(semver.Maybe(c.Version, c.Release, c.Mono), true)
	if err != nil {
		return fmt.Errorf("cannot uninstall godot: %w", err)
	}

	return nil
}

type Install struct {
	Version                string `arg:"" help:"Godot engine version to download and install in the format x.x.x.x, x.x.x or x.x"`
	ExcludeExportTemplates bool   `short:"e" help:"Exclude export templates in install"`
	Release                string `short:"r" default:"stable" help:"Release to use (dev1, alpha2, beta3, rc4, stable, etc)"`
	Mono                   bool   `short:"m" help:"Use mono version"`
}

func (c *Install) Run(app *gevm.App) error {
	if !c.ExcludeExportTemplates {
		err := app.ExportTemplates.Install(semver.Maybe(c.Version, c.Release, c.Mono))
		if err != nil {
			return fmt.Errorf("cannot install export templates: %w", err)
		}
	}

	err := app.Godot.Install(semver.Maybe(c.Version, c.Release, c.Mono))
	if err != nil {
		return fmt.Errorf("cannot install godot: %w", err)
	}

	return nil
}

type Path struct {
	Version string `arg:"" help:"Godot engine version to use in the format x.x.x.x, x.x.x or x.x"`
	Release string `short:"r" default:"stable" help:"Release to use (dev1, alpha2, beta3, rc4, stable, etc)"`
	Mono    bool   `short:"m" help:"Use mono version"`
}

func (c *Path) Run(app *gevm.App) error {
	err := app.Godot.Path(semver.Maybe(c.Version, c.Release, c.Mono))
	if err != nil {
		return fmt.Errorf("cannot print path: %w", err)
	}

	return nil
}

type List struct{}

func (c *List) Run(app *gevm.App) error {
	err := app.Godot.List()
	if err != nil {
		return fmt.Errorf("cannot list godot: %w", err)
	}

	return nil
}

type Clear struct {
	ExcludeExportTemplates bool `short:"e" help:"Exclude export templates in uninstall"`
}

func (c *Clear) Run(app *gevm.App) error {
	if !c.ExcludeExportTemplates {
		err := app.ExportTemplates.Clear()
		if err != nil {
			return fmt.Errorf("cannot clear export templates: %w", err)
		}
	}

	err := app.Godot.Clear()
	if err != nil {
		return fmt.Errorf("cannot clear godot: %w", err)
	}

	return nil
}

type Godot struct {
	Download  Download  `cmd:"" help:"Download godot engine to the cache by version"`
	Uninstall Uninstall `cmd:"" help:"Uninstall godot engine by version"`
	Install   Install   `cmd:"" help:"Install godot engine by version"`
	Path      Path      `cmd:"" help:"Print path to godot engine version"`
	List      List      `cmd:"" help:"List all current godot engine versions"`
	Clear     Clear     `cmd:"" help:"Clear all godot engine versions"`
}
