package pm

import "fmt"

type PrinterList []Printer

func MergePrinter(printers ...Printer) Printer {
	p := PrinterList(printers)

	return p
}

func (printers PrinterList) Prints() (DevicePrinters, error) {
	allPrinters := DevicePrinters{}
	for _, printer := range printers {
		prints, err := printer.Prints()
		if err != nil {
			return nil, fmt.Errorf("error generating prints: %v", err)
		}
		for _, p := range prints {
			allPrinters = append(allPrinters, p)
		}
	}

	return allPrinters, nil
}
