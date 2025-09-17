package resource

import (
	koordletuti "github.com/koordinator-sh/koordinator/pkg/koordlet/util"
)

type npuDevice struct {
	DeviceInfo koordletuti.XPUDeviceInfo
}

func (device npuDevice) GetDeviceInfo() (koordletuti.XPUDeviceInfo, error) {
	return device.DeviceInfo, nil
}
