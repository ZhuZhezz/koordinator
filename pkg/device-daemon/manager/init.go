package pm

import (
	"github.com/koordinator-sh/koordinator/pkg/device-daemon/resource"
)

var PrinterMap map[string]func(manager resource.Manager) (Printer, error)

func init() {
	PrinterMap = make(map[string]func(manager resource.Manager) (Printer, error))
	PrinterMap[MLU] = newMLUPrinter
	PrinterMap[NPU] = newNPUPrinter
	PrinterMap[XPU] = newXPUPrinter
}
