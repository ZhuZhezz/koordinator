package pm

import (
	"fmt"

	resourceconifg "github.com/koordinator-sh/koordinator/cmd/koord-device-daemon/config/v1"
	"github.com/koordinator-sh/koordinator/pkg/device-daemon/resource"
)

func NewDevicePrinter(manager map[string]resource.Manager, config *resourceconifg.Config) (Printer, error) {
	printers, err := getPrinters(manager)
	if err != nil {
		return nil, fmt.Errorf("failed to get printers: %v", err)
	}

	return printers, nil
}

func getPrinters(manager map[string]resource.Manager) (Printer, error) {
	var printers PrinterList
	for name, manager := range manager {
		if printerFunc, ok := PrinterMap[name]; ok {
			printer, err := printerFunc(manager)
			if err != nil {
				return nil, fmt.Errorf("error creating printer: %v", err)
			}
			printers = append(printers, printer)
		}
	}
	return printers, nil
}
