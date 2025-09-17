package pm

import (
	koordletuti "github.com/koordinator-sh/koordinator/pkg/koordlet/util"
)

type DevicePrinters koordletuti.XPUDevices

func (printers DevicePrinters) Prints() (DevicePrinters, error) {
	return printers, nil
}
