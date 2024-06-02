package mac

import (
	"context"
	"devtool/internal/hardware/devices/keyboard"
	"devtool/internal/hardware/devices/mouse"
	"devtool/internal/hardware/devices/other"
	iscanner "devtool/internal/hardware/scanner_interface"
	"sync"
)

type macScanner struct {
}

func NewDeviceScanner() (*macScanner, error) {
	return &macScanner{}, nil
}

func (s *macScanner) GetDevices(ctx context.Context) ([]iscanner.Device, error) {
	devCh := make(chan iscanner.Device, 10)
	errCh := make(chan error)
	defer close(errCh)
	wg := &sync.WaitGroup{}

	runners := []func(ctx context.Context, wg *sync.WaitGroup, devChan chan<- iscanner.Device, errChan chan<- error){
		other.GetDevices,
		mouse.GetDevices,
		keyboard.GetDevices,
	}

	wg.Add(len(runners))
	for _, runner := range runners {
		go runner(ctx, wg, devCh, errCh)
	}

	go func() {
		wg.Wait()
		close(devCh)
	}()

	var devices []iscanner.Device
	for {
		select {
		case <-ctx.Done():
			return devices, ctx.Err()
		case err := <-errCh:
			return devices, err

		case dev, ok := <-devCh:
			if ok {
				devices = append(devices, dev)
			} else {
				return devices, nil
			}
		}
	}
}
