package devices

import (
	"devtool/internal/hardware"
	iscanner "devtool/internal/hardware/scanner_interface"
	"devtool/proto_generated"
	"google.golang.org/protobuf/proto"
	"net/http"
)

type DevicesHandler interface {
	GetConnectedDevices(w http.ResponseWriter, r *http.Request)
}

type devicesHandler struct {
	scanner iscanner.DeviceScanner
}

func NewDevicesHandler() (DevicesHandler, error) {
	scanner, err := hardware.NewDeviceScanner()
	if err != nil {
		return nil, err
	}
	return &devicesHandler{scanner: scanner}, nil
}

func (handler *devicesHandler) GetConnectedDevices(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	getConnectedDevicesOperationDone := make(chan struct{})
	var deviceList []iscanner.Device
	var deviceErr error

	go func() {
		deviceList, deviceErr = handler.scanner.GetDevices(ctx)
		close(getConnectedDevicesOperationDone)
	}()

	select {
	case <-ctx.Done():
		http.Error(w, ctx.Err().Error(), http.StatusInternalServerError)
		return

	case <-getConnectedDevicesOperationDone:
		if deviceErr != nil {
			http.Error(w, deviceErr.Error(), http.StatusInternalServerError)
			return
		}

		protoDeviceList := &proto_generated.DeviceList{
			Devices: []*proto_generated.Device{},
		}

		for _, device := range deviceList {
			protoDeviceList.Devices = append(protoDeviceList.Devices, &proto_generated.Device{
				Type: device.Type,
				Path: device.Path,
				Vid:  device.VID,
				Pid:  device.PID,
			})
		}

		serialized, err := proto.Marshal(protoDeviceList)
		if err != nil {
			http.Error(w, "Failed to marshal device list", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/x-protobuf")
		w.Write(serialized)
	}
}
