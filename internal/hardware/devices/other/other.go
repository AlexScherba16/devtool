package other

import (
	"context"
	iscanner "devtool/internal/hardware/scanner_interface"
	"devtool/pkg/hardware_interface/usb"
	"errors"
	"fmt"
	"github.com/google/gousb"
	"sync"
)

func GetDevices(ctx context.Context, wg *sync.WaitGroup, devChan chan<- iscanner.Device, errChan chan<- error) {
	usbCtx := gousb.NewContext()
	defer usbCtx.Close()
	defer wg.Done()

	devices, err := usbCtx.OpenDevices(usb.IsOther)
	if err != nil {
		message := fmt.Sprintf("Error occured while opening other devices: %v", err)
		errChan <- errors.New(message)
	}

	defer func() {
		for _, d := range devices {
			d.Close()
		}
	}()

	for _, dev := range devices {
		select {
		case <-ctx.Done():
			return
		default:
			devChan <- iscanner.Device{
				Type: "Other",
				Path: usb.PortsToPath(dev.Desc),
				VID:  dev.Desc.Vendor.String(),
				PID:  dev.Desc.Product.String(),
			}
		}
	}
}
