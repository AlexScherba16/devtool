package os

import "runtime"

const (
	MAC        = "darvin"
	LINUX      = "linux"
	WINDOWS    = "windows"
	UNDETECTED = "undetected"
)

type HostOs string

func GetHostOS() HostOs {
	switch runtime.GOOS {
	case "darwin":
		return MAC
	case "linux":
		return LINUX
	case "windows":
		return WINDOWS
	default:
		return UNDETECTED
	}
}
