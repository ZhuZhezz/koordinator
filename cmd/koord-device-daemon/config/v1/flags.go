package v1

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// prt returns a reference to whatever type is passed into it
func ptr[T any](x T) *T {
	return &x
}

// updateFromCLIFlag conditionally updates the config flag at 'pflag' to the value of the CLI flag with name 'flagName'
func updateFromCLIFlag[T any](pflag **T, c *cli.Context, flagName string) {
	if c.IsSet(flagName) || *pflag == (*T)(nil) {
		switch flag := any(pflag).(type) {
		case **string:
			*flag = ptr(c.String(flagName))
		case **[]string:
			*flag = ptr(c.StringSlice(flagName))
		case **bool:
			*flag = ptr(c.Bool(flagName))
		case **Duration:
			*flag = ptr(Duration(c.Duration(flagName)))
		default:
			panic(fmt.Errorf("unsupported flag type for %v: %T", flagName, flag))
		}
	}
}

// Flags holds the full list of flags used to configure the device plugin and KDD.
type Flags struct {
	CommandLineFlags
}

// CommandLineFlags holds the list of command line flags used to configure the device plugin and KDD.
type CommandLineFlags struct {
	KDD *KDDCommandLineFlags `json:"kdd,omitempty" yaml:"kdd,omitempty"`
}

// KDDCommandLineFlags holds the list of command line flags specific to KDD.
type KDDCommandLineFlags struct {
	Oneshot          *bool     `json:"oneshot"         yaml:"oneshot"`
	SleepInterval    *Duration `json:"sleepInterval"   yaml:"sleepInterval"`
	PrintsOutputFile *string   `json:"printsOutputFile" yaml:"printsOutputFile"`
}

// UpdateFromCLIFlags updates Flags from settings in the cli Flags if they are set.
func (f *Flags) UpdateFromCLIFlags(c *cli.Context, flags []cli.Flag) {
	for _, flag := range flags {
		for _, n := range flag.Names() {
			// KDD specific flags
			if f.KDD == nil {
				f.KDD = &KDDCommandLineFlags{}
			}
			switch n {
			case "oneshot":
				updateFromCLIFlag(&f.KDD.Oneshot, c, n)
			case "prints-output-file":
				updateFromCLIFlag(&f.KDD.PrintsOutputFile, c, n)
			case "sleep-interval":
				updateFromCLIFlag(&f.KDD.SleepInterval, c, n)
			}
		}
	}
}
