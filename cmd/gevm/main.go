package main

import (
	"log"

	"github.com/alecthomas/kong"
	"github.com/bashidogames/gevm"
	"github.com/bashidogames/gevm/cmd/gevm/cache"
	"github.com/bashidogames/gevm/cmd/gevm/exporttemplates"
	"github.com/bashidogames/gevm/cmd/gevm/godot"
	"github.com/bashidogames/gevm/cmd/gevm/settings"
	"github.com/bashidogames/gevm/cmd/gevm/shortcuts"
	"github.com/bashidogames/gevm/cmd/gevm/version"
	"github.com/bashidogames/gevm/cmd/gevm/versions"
	"github.com/bashidogames/gevm/config"
	"github.com/bashidogames/gevm/internal/logging"
)

var CLI struct {
	Versions        versions.Versions               `cmd:"" help:"View available versions for download"`
	ExportTemplates exporttemplates.ExportTemplates `cmd:"" help:"Run commands related to export templates"`
	Godot           godot.Godot                     `cmd:"" help:"Run commands related to godot engines"`
	Shortcuts       shortcuts.Shortcuts             `cmd:"" help:"Remove and add godot shortcuts"`
	Settings        settings.Settings               `cmd:"" help:"View and adjust config settings"`
	Cache           cache.Cache                     `cmd:"" help:"Run commands on the cache"`
	Version         version.Version                 `cmd:"" help:"Print current version"`

	LoggingLevel string `short:"l" enum:"nothing,error,warning,info,debug,trace" default:"info" help:"Which log level to use"`
	ConfigPath   string `help:"Override which config path to use"`
	Silent       bool   `help:"Prevent progress bar log spam"`
}

func main() {
	ctx := kong.Parse(&CLI)

	var level logging.Level
	switch CLI.LoggingLevel {
	default:
		log.Fatalf("logging level not handled: %s", CLI.LoggingLevel)
	case "nothing":
		level = logging.NOTHING
	case "error":
		level = logging.ERROR
	case "warning":
		level = logging.WARNING
	case "info":
		level = logging.INFO
	case "debug":
		level = logging.DEBUG
	case "trace":
		level = logging.TRACE
	}

	logger, err := logging.New(level)
	if err != nil {
		log.Fatalf("failed to create logger: %s", err)
	}

	config, err := config.New(
		config.OptionSetConfigPath(CLI.ConfigPath),
		config.OptionSetSilent(CLI.Silent),
		config.OptionSetLogger(logger),
	)
	if err != nil {
		log.Fatalf("failed to create config: %s", err)
	}

	app, err := gevm.New(config)
	if err != nil {
		log.Fatalf("failed to create app: %s", err)
	}

	err = ctx.Run(app)
	if err != nil {
		log.Fatalf("failed to run: %s", err)
	}
}
