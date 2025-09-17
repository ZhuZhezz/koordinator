package resource

import (
	koordletuti "github.com/koordinator-sh/koordinator/pkg/koordlet/util"
)

type mluDevice struct {
	DeviceInfo koordletuti.XPUDeviceInfo
}

func (device mluDevice) GetDeviceInfo() (koordletuti.XPUDeviceInfo, error) {
	return device.DeviceInfo, nil
}
