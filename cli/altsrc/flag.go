package altsrc

import (
	"github.com/rancher/spur/cli"
)

// FlagInputSourceExtension is an extension interface of cli.Flag that
// allows a value to be set on the existing parsed flags.
type FlagInputSourceExtension interface {
	cli.Flag
	ApplyInputSourceValue(context *cli.Context, isc cli.InputSourceContext) error
}

// ApplyInputSourceValues iterates over all provided flags and
// executes ApplyInputSourceValue on flags implementing the
// FlagInputSourceExtension interface to initialize these flags
// to an alternate input source.
func ApplyInputSourceValues(context *cli.Context, inputSourceContext cli.InputSourceContext, flags []cli.Flag) error {
	for _, f := range flags {
		inputSourceExtendedFlag, isType := f.(FlagInputSourceExtension)
		if isType {
			err := inputSourceExtendedFlag.ApplyInputSourceValue(context, inputSourceContext)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// InitInputSource is used to to setup an cli.InputSourceContext on a cli.Command Before method. It will create a new
// input source based on the func provided. If there is no error it will then apply the new input source to any flags
// that are supported by the input source
func InitInputSource(flags []cli.Flag, createInputSource func() (cli.InputSourceContext, error)) cli.BeforeFunc {
	return func(context *cli.Context) error {
		inputSource, err := createInputSource()
		if err != nil {
			return err
		}

		return ApplyInputSourceValues(context, inputSource, flags)
	}
}

// InitInputSourceWithContext is used to to setup an cli.InputSourceContext on a cli.Command Before method. It will create a new
// input source based on the func provided with potentially using existing cli.Context values to initialize itself. If there is
// no error it will then apply the new input source to any flags that are supported by the input source
func InitInputSourceWithContext(flags []cli.Flag, createInputSource func(context *cli.Context) (cli.InputSourceContext, error)) cli.BeforeFunc {
	return func(context *cli.Context) error {
		inputSource, err := createInputSource(context)
		if err != nil {
			return err
		}

		return ApplyInputSourceValues(context, inputSource, flags)
	}
}
