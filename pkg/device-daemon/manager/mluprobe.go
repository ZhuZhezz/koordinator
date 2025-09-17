package pm

import (
	"fmt"

	"k8s.io/klog/v2"

	"github.com/koordinator-sh/koordinator/pkg/device-daemon/resource"
)

func newMLUPrinter(manager resource.Manager) (Printer, error) {
	klog.Info("Creating MLU printer")
	mluDevices, err := manager.GetDevices()
	if err != nil {
		return nil, fmt.Errorf("error getting mlu devices: %v", err)
	}

	mluCount := len(mluDevices)
	if mluCount == 0 {
		return DevicePrinters{}, nil
	}

	mluPrinters := DevicePrinters{}
	for _, mlu := range mluDevices {
		deviceInfo, err := mlu.GetDeviceInfo()
		if err != nil {
			return nil, fmt.Errorf("error getting mlu device info: %v", err)
		}
		mluPrinters = append(mluPrinters, deviceInfo)
	}

	return mluPrinters, nil
}
