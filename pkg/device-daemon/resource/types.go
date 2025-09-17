package resource

import (
	koordletuti "github.com/koordinator-sh/koordinator/pkg/koordlet/util"
)

// Manager defines an interface for managing devices
//
//go:generate moq -rm -out manager_mock.go . Manager
type Manager interface {
	GetDevices() ([]Device, error)
	GetDriverVersion() (string, error)
}

// Device defines an interface for a device with which labels are associated
//
//go:generate moq -out device_mock.go . Device
type Device interface {
	GetDeviceInfo() (koordletuti.XPUDeviceInfo, error)
}
