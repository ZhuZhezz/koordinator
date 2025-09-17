package resource

import (
	koordletuti "github.com/koordinator-sh/koordinator/pkg/koordlet/util"
)

type xpuDevice struct {
	DeviceInfo koordletuti.XPUDeviceInfo
}

func (device xpuDevice) GetDeviceInfo() (koordletuti.XPUDeviceInfo, error) {
	return device.DeviceInfo, nil
}
