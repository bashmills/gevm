package locator

import (
	"path/filepath"
	"regexp"

	"github.com/bashidogames/gevm/config"
	"github.com/bashidogames/gevm/internal/utils"
	"github.com/bashidogames/gevm/semver"
)

const EXECUTABLE_REGEX_PATTERN = "Godot(.*?)[.]exe"
const INVALID_REGEX_PATTERN = "(console)[.]exe"

var ExecutableRegex = regexp.MustCompile(EXECUTABLE_REGEX_PATTERN)
var InvalidRegex = regexp.MustCompile(INVALID_REGEX_PATTERN)

type Locator struct {
	Config *config.Config
}

func (l *Locator) TargetPath(semver semver.Semver) (string, error) {
	return l.locateExecutable(filepath.Join(l.Config.GodotRootDirectory, semver.GodotString()))
}

func (l *Locator) locateExecutable(root string) (string, error) {
	return utils.LocateExecutable(func(filename string) bool {
		return ExecutableRegex.MatchString(filename) && !InvalidRegex.MatchString(filename)
	}, root, false)
}

func New(config *config.Config) (*Locator, error) {
	return &Locator{
		Config: config,
	}, nil
}
