package cli

import (
	"time"

	"github.com/rancher/spur/flag"
)

var _ = time.Time{}

// BoolSlice is a type alias for []bool
type BoolSlice = []bool

// BoolSliceFlag is a flag with type []bool
type BoolSliceFlag struct {
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

	Value       BoolSlice
	Destination *BoolSlice
}

// Apply populates the flag given the flag set and environment
func (f *BoolSliceFlag) Apply(set *flag.FlagSet) error {
	return Apply(f, "bool slice", set)
}

// ApplyInputSourceValue applies a BoolSlice value to the flagSet if required
func (f *BoolSliceFlag) ApplyInputSourceValue(context *Context, isc InputSourceContext) error {
	return ApplyInputSourceValue(f, context, isc)
}

// String returns a readable representation of this value
// (for usage defaults)
func (f *BoolSliceFlag) String() string {
	return FlagStringer(f)
}

// BoolSlice looks up the value of a local BoolSliceFlag, returns
// an empty value if not found
func (c *Context) BoolSlice(name string) []bool {
	return c.Lookup(name, *new(BoolSlice)).([]bool)
}
