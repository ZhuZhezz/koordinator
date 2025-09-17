package resource

import (
	"fmt"
	"strconv"

	"github.com/jaypipes/ghw"
	"k8s.io/klog/v2"

	koordletuti "github.com/koordinator-sh/koordinator/pkg/koordlet/util"
)

type xpuManager struct{}

func NewXPUManager() Manager {
	return &xpuManager{}
}

func (manager *xpuManager) GetDevices() ([]Device, error) {
	klog.Info("getXPUDevice(): Getting XPU devices")
	pci, err := ghw.PCI()
	if err != nil {
		return nil, fmt.Errorf("getXPUDevice(): new PCI instance error, %v", err)
	}
	devices := pci.ListDevices()
	if len(devices) == 0 {
		klog.Info("getXPUDevice(): no pci devices")
		return nil, nil
	}
	xpuExists := false
	for _, device := range devices {
		if IsXPUDevice(device) {
			xpuExists = true
			break
		}
	}
	if !xpuExists {
		return []Device{}, nil
	}
	var xpuDevices []Device

	count, err := GetXPUCount(XPU)
	if err != nil {
		return nil, fmt.Errorf("getXPUDevice(): get GPU count, %v", err)
	}

	productName, err := GetXPUName(XPU)
	if err != nil {
		return nil, fmt.Errorf("getXPUDevice(): get product name error, %v", err)
	}

	mem, err := GetXPUMemory(XPU)
	if err != nil {
		return nil, fmt.Errorf("getXPUDevice(): get memory error, %v", err)
	}

	for i := 0; i < count; i++ {
		topo, status, uuid, err := GetDeviceInfo(XPU, fmt.Sprintf("%d", i))
		if err != nil {
			klog.Errorf("getXPUDevice(): get device info error, %v", err)
			continue
		}
		deviceInfo := koordletuti.XPUDeviceInfo{
			Vendor: "kunlun",
			Model:  productName,
			UUID:   uuid,
			Minor:  strconv.Itoa(i),
			Resources: map[string]string{
				"koordinator.sh/gpu-memory": mem,
				"koordinator.sh/gpu-core":   "100",
			},
			Topology: &topo,
			Status:   &status,
		}
		xpuDevice := xpuDevice{
			DeviceInfo: deviceInfo,
		}
		xpuDevices = append(xpuDevices, xpuDevice)
	}

	klog.Infof("gpu devices: %+v", xpuDevices)
	return xpuDevices, nil
}

func (manager *xpuManager) GetDriverVersion() (string, error) {
	return "", nil
}
