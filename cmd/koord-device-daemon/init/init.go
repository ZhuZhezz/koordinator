package init

import (
	"github.com/koordinator-sh/koordinator/pkg/device-daemon/resource"
)

var ManagerMap map[string]resource.Manager

const (
	NPU = "npu"
	MLU = "mlu"
	XPU = "xpu"
)

func init() {
	ManagerMap = make(map[string]resource.Manager)
	ManagerMap[NPU] = resource.NewNPUManager()
	ManagerMap[MLU] = resource.NewMLUManager()
	ManagerMap[XPU] = resource.NewXPUManager()
}
