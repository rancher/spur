package cli

import (
	"time"

	"github.com/rancher/spur/flag"
)

var _ = time.Time{}

// DurationSlice is a type alias for []time.Duration
type DurationSlice = []time.Duration

// DurationSliceFlag is a flag with type []time.Duration
type DurationSliceFlag struct {
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

	Value       DurationSlice
	Destination *DurationSlice
}

// Apply populates the flag given the flag set and environment
func (f *DurationSliceFlag) Apply(set *flag.FlagSet) error {
	return Apply(f, "duration slice", set)
}

// ApplyInputSourceValue applies a DurationSlice value to the flagSet if required
func (f *DurationSliceFlag) ApplyInputSourceValue(context *Context, isc InputSourceContext) error {
	return ApplyInputSourceValue(f, context, isc)
}

// String returns a readable representation of this value
// (for usage defaults)
func (f *DurationSliceFlag) String() string {
	return FlagStringer(f)
}

// DurationSlice looks up the value of a local DurationSliceFlag, returns
// an empty value if not found
func (c *Context) DurationSlice(name string) []time.Duration {
	return c.Lookup(name, *new(DurationSlice)).([]time.Duration)
}
