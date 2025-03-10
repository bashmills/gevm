package locator

import (
	"fmt"
	"path/filepath"
	"regexp"

	"github.com/bashidogames/gevm/config"
	"github.com/bashidogames/gevm/internal/utils"
	"github.com/bashidogames/gevm/semver"
)

const EXECUTABLE_REGEX_PATTERN = "Godot(.*?)[.]exe"
const INVALID_REGEX_PATTERN = "(console)[.]exe"
const SHORTCUT_FILENAME = "Godot %s.lnk"
const SHORTCUT_NAME = "Godot %s"
const LINK_FILENAME = "godot.exe"

var ExecutableRegex = regexp.MustCompile(EXECUTABLE_REGEX_PATTERN)
var InvalidRegex = regexp.MustCompile(INVALID_REGEX_PATTERN)

type Locator struct {
	Config *config.Config
}

func (l *Locator) ApplicationShortcutPath(semver semver.Semver) string {
	return filepath.Join(l.Config.ApplicationShortcutDirectory, l.shortcutFilename(semver))
}

func (l *Locator) DesktopShortcutPath(semver semver.Semver) string {
	return filepath.Join(l.Config.DesktopShortcutDirectory, l.shortcutFilename(semver))
}

func (l *Locator) TargetPath(semver semver.Semver) (string, error) {
	return l.locateExecutable(filepath.Join(l.Config.GodotRootDirectory, semver.GodotString()))
}

func (l *Locator) ShortcutName(semver semver.Semver) string {
	return fmt.Sprintf(SHORTCUT_NAME, semver.GodotString())
}

func (l *Locator) LinkPath(semver semver.Semver) string {
	return filepath.Join(l.Config.BinDirectory, LINK_FILENAME)
}

func (l *Locator) shortcutFilename(semver semver.Semver) string {
	return fmt.Sprintf(SHORTCUT_FILENAME, semver.GodotString())
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
