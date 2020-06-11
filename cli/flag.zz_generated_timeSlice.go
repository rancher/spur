package cli

import (
	"time"

	"github.com/rancher/spur/flag"
)

var _ = time.Time{}

// TimeSlice is a type alias for []time.Time
type TimeSlice = []time.Time

// TimeSliceFlag is a flag with type []time.Time
type TimeSliceFlag struct {
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

	Value       TimeSlice
	Destination *TimeSlice
}

// Apply populates the flag given the flag set and environment
func (f *TimeSliceFlag) Apply(set *flag.FlagSet) error {
	return Apply(f, "time slice", set)
}

// ApplyInputSourceValue applies a TimeSlice value to the flagSet if required
func (f *TimeSliceFlag) ApplyInputSourceValue(context *Context, isc InputSourceContext) error {
	return ApplyInputSourceValue(f, context, isc)
}

// String returns a readable representation of this value
// (for usage defaults)
func (f *TimeSliceFlag) String() string {
	return FlagStringer(f)
}

// TimeSlice looks up the value of a local TimeSliceFlag, returns
// an empty value if not found
func (c *Context) TimeSlice(name string) []time.Time {
	return c.Lookup(name, *new(TimeSlice)).([]time.Time)
}
