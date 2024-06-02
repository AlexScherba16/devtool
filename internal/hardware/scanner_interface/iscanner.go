package scanner_interface

import "context"

type Device struct {
	Type string
	Path string
	VID  string
	PID  string
}

type DeviceScanner interface {
	GetDevices(ctx context.Context) ([]Device, error)
}
