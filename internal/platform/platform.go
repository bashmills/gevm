package platform

import (
	"fmt"
	"runtime"
)

type Platform string

const (
	ExportTemplates Platform = "Export Templates"
	WindowsArm64    Platform = "Windows Arm64"
	WindowsAmd64    Platform = "Windows Amd64"
	DarwinArm64     Platform = "Darwin Arm64"
	DarwinAmd64     Platform = "Darwin Amd64"
	LinuxArm64      Platform = "Linux Arm64"
	LinuxAmd64      Platform = "Linux Amd64"
)

var Platforms = []Platform{
	ExportTemplates,
	WindowsArm64,
	WindowsAmd64,
	DarwinArm64,
	DarwinAmd64,
	LinuxArm64,
	LinuxAmd64,
}

func Get() (Platform, error) {
	switch runtime.GOOS {
	case "windows":
		switch runtime.GOARCH {
		case "arm64":
			return WindowsArm64, nil
		case "amd64":
			return WindowsAmd64, nil
		}
	case "darwin":
		switch runtime.GOARCH {
		case "arm64":
			return DarwinArm64, nil
		case "amd64":
			return DarwinAmd64, nil
		}
	case "linux":
		switch runtime.GOARCH {
		case "arm64":
			return LinuxArm64, nil
		case "amd64":
			return LinuxAmd64, nil
		}
	}

	return "", fmt.Errorf("invalid platform")
}
