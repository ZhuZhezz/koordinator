package pm

import (
	"fmt"

	"k8s.io/klog/v2"

	"github.com/koordinator-sh/koordinator/pkg/device-daemon/resource"
)

func newXPUPrinter(manager resource.Manager) (Printer, error) {
	klog.Info("Creating XPU printer")
	xpuDevices, err := manager.GetDevices()
	if err != nil {
		return nil, fmt.Errorf("error getting XPU devices: %v", err)
	}

	xpuCount := len(xpuDevices)
	if xpuCount == 0 {
		return DevicePrinters{}, nil
	}

	xpuPrinters := DevicePrinters{}
	for _, xpu := range xpuDevices {
		deviceInfo, err := xpu.GetDeviceInfo()
		if err != nil {
			return nil, fmt.Errorf("error getting XPU device info: %v", err)
		}
		xpuPrinters = append(xpuPrinters, deviceInfo)
	}

	return xpuPrinters, nil
}
