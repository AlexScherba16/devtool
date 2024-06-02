package hardware

import (
	"devtool/internal/hardware/os/mac"
	iscanner "devtool/internal/hardware/scanner_interface"
	"devtool/pkg/os"
	"fmt"
)

func NewDeviceScanner() (iscanner.DeviceScanner, error) {
	hostOs := os.GetHostOS()

	switch hostOs {
	case os.MAC:
		return mac.NewDeviceScanner()
	default:
		return nil, fmt.Errorf("unsupported OS: %s", hostOs)
	}
}
