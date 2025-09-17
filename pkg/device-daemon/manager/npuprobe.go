package pm

import (
	"fmt"

	"k8s.io/klog/v2"

	"github.com/koordinator-sh/koordinator/pkg/device-daemon/resource"
)

func newNPUPrinter(manager resource.Manager) (Printer, error) {
	klog.Info("Creating NPU printer")
	npuDevices, err := manager.GetDevices()
	if err != nil {
		return nil, fmt.Errorf("error getting npu devices: %v", err)
	}

	npuCount := len(npuDevices)
	if npuCount == 0 {
		return DevicePrinters{}, nil
	}

	npuPrinters := DevicePrinters{}
	for _, npu := range npuDevices {
		deviceInfo, err := npu.GetDeviceInfo()
		if err != nil {
			return nil, fmt.Errorf("error getting npu device info: %v", err)
		}
		npuPrinters = append(npuPrinters, deviceInfo)
	}

	return npuPrinters, nil
}
