package cli

import (
	"time"

	"github.com/rancher/spur/flag"
)

var _ = time.Time{}

// UintSlice is a type alias for []uint
type UintSlice = []uint

// UintSliceFlag is a flag with type []uint
type UintSliceFlag struct {
	Name        string
	Aliases     []string
	EnvVars     []string
	Usage       string
	DefaultText string
	FilePath    string
	Required    bool
	Hidden      bool
	TakesFile   bool
	SkipAltSrc  bool
	LoadedValue bool

	Value       UintSlice
	Destination *UintSlice
}

// Apply populates the flag given the flag set and environment
func (f *UintSliceFlag) Apply(set *flag.FlagSet) error {
	return Apply(f, "uint slice", set)
}

// ApplyInputSourceValue applies a UintSlice value to the flagSet if required
func (f *UintSliceFlag) ApplyInputSourceValue(context *Context, isc InputSourceContext) error {
	return ApplyInputSourceValue(f, context, isc)
}

// String returns a readable representation of this value
// (for usage defaults)
func (f *UintSliceFlag) String() string {
	return FlagStringer(f)
}

// UintSlice looks up the value of a local UintSliceFlag, returns
// an empty value if not found
func (c *Context) UintSlice(name string) []uint {
	return c.Lookup(name, *new(UintSlice)).([]uint)
}
