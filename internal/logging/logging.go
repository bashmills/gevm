package logging

import (
	"fmt"

	"github.com/bashmills/gevm/internal/utils"
	"github.com/bashmills/gevm/logger"
)

type Level int64

const (
	NOTHING Level = iota
	ERROR
	WARNING
	INFO
	DEBUG
	TRACE
)

type Logging struct {
	Level Level
}

func (l Logging) Error(format string, a ...any) {
	if !l.ShouldLog(ERROR) {
		return
	}

	utils.Printlnf("ERROR: %s", fmt.Sprintf(format, a...))
}

func (l Logging) Warning(format string, a ...any) {
	if !l.ShouldLog(WARNING) {
		return
	}

	utils.Printlnf("WARNING: %s", fmt.Sprintf(format, a...))
}

func (l Logging) Info(format string, a ...any) {
	if !l.ShouldLog(INFO) {
		return
	}

	utils.Printlnf(format, a...)
}

func (l Logging) Debug(format string, a ...any) {
	if !l.ShouldLog(DEBUG) {
		return
	}

	utils.Printlnf(format, a...)
}

func (l Logging) Trace(format string, a ...any) {
	if !l.ShouldLog(TRACE) {
		return
	}

	utils.Printlnf(format, a...)
}

func (l Logging) ShouldLog(level Level) bool {
	return l.Level >= level
}

func New(level Level) (logger.Logger, error) {
	return &Logging{
		Level: level,
	}, nil
}
