package pm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/google/renameio"
	"k8s.io/klog/v2"

	resourceconifg "github.com/koordinator-sh/koordinator/cmd/koord-device-daemon/config/v1"
)

// Outputer defines a mechanism to output labels and device infos.
type Outputer interface {
	OutputPrints(DevicePrinters) error
}

func NewPrintsOutputer(config *resourceconifg.Config) (Outputer, error) {

	return ToFile(*config.Flags.KDD.PrintsOutputFile), nil
}

func ToFile(path string) Outputer {
	if path == "" {
		return &toWriter{os.Stdout}
	}

	o := toFile(path)
	return &o
}

// toFile writes to the specified file.
type toFile string

type toWriter struct {
	io.Writer
}

func (output *toWriter) OutputPrints(printers DevicePrinters) error {
	return nil
}

func (path *toFile) OutputPrints(printers DevicePrinters) error {
	klog.Infof("Writing device printers (JSON) to output file %v", *path)

	data, err := json.MarshalIndent(printers, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling device printers to JSON: %v", err)
	}

	buffer := bytes.NewBuffer(data)
	if err := renameio.WriteFile(string(*path), buffer.Bytes(), 0644); err != nil {
		return fmt.Errorf("error atomically writing file '%s': %w", *path, err)
	}
	return nil
}
