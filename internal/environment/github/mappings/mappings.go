package mappings

import "github.com/bashmills/gevm/internal/platform"

type Mapping struct {
	System []string
	Arch   []string
}

var Mappings = map[platform.Platform]Mapping{
	platform.ExportTemplates: {
		System: []string{"export"},
		Arch:   []string{"templates"},
	},
	platform.WindowsArm64: {
		System: []string{"windows", "win"},
		Arch:   []string{"arm64"},
	},
	platform.WindowsAmd64: {
		System: []string{"windows", "win"},
		Arch:   []string{"64"},
	},
	platform.DarwinArm64: {
		System: []string{"macos", "osx"},
		Arch:   []string{"universal"},
	},
	platform.DarwinAmd64: {
		System: []string{"macos", "osx"},
		Arch:   []string{"universal", "fat", "64"},
	},
	platform.LinuxArm64: {
		System: []string{"linux", "x11"},
		Arch:   []string{"arm64"},
	},
	platform.LinuxAmd64: {
		System: []string{"linux", "x11"},
		Arch:   []string{"64"},
	},
}

var Overrides = map[platform.Platform][]string{
	platform.DarwinAmd64: {"universal"},
}
