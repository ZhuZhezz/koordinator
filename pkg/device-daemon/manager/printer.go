package pm

import (
	"fmt"

	resourceconifg "github.com/koordinator-sh/koordinator/cmd/koord-device-daemon/config/v1"
	"github.com/koordinator-sh/koordinator/pkg/device-daemon/resource"
)

// Printer defines an interface for generating device printers
type Printer interface {
	Prints() (DevicePrinters, error)
}

// NewPrinter constructs the required printer from the specified config
func NewPrinter(manager map[string]resource.Manager, config *resourceconifg.Config) (Printer, error) {
	devicePrinter, err := NewDevicePrinter(manager, config)
	if err != nil {
		return nil, fmt.Errorf("error creating devicePrinter: %v", err)
	}

	return devicePrinter, nil
}
