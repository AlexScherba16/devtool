package usb

import (
	"fmt"
	"github.com/google/gousb"
	"strings"
)

// https://docs.silabs.com/protocol-usb/1.2.0/protocol-usb-api/usbd-hid
const (
	// http://www.on-time.com/rtos-32-docs/rtusb-32/programming-manual/programming-with/class-drivers/keyboards.htm
	USBD_HID_PROTOCOL_KBD = 0x01

	// http://www.on-time.com/rtos-32-docs/rtusb-32/programming-manual/programming-with/class-drivers/mice.htm
	USBD_HID_PROTOCOL_MOUSE = 0x02
)

func IsMouse(desc *gousb.DeviceDesc) bool {
	if desc == nil {
		return false
	}
	if desc.Class == gousb.ClassHID && desc.Protocol == USBD_HID_PROTOCOL_MOUSE {
		return true
	}
	return false
}

func IsKeyboard(desc *gousb.DeviceDesc) bool {
	if desc == nil {
		return false
	}
	if desc.Class == gousb.ClassHID && desc.Protocol == USBD_HID_PROTOCOL_KBD {
		return true
	}
	return false
}

func IsOther(desc *gousb.DeviceDesc) bool {
	if desc == nil {
		return false
	}
	if desc.Class != gousb.ClassHID && desc.Protocol != USBD_HID_PROTOCOL_KBD && desc.Protocol != USBD_HID_PROTOCOL_MOUSE {
		return true
	}
	return false
}

func PortsToPath(desc *gousb.DeviceDesc) string {
	if desc == nil {
		return ""
	}

	ports := make([]string, len(desc.Path))
	for i, p := range desc.Path {
		ports[i] = fmt.Sprintf("%d", p)
	}
	return strings.Join(ports, "/")
}
