package cli

import (
	"time"

	"github.com/rancher/spur/flag"
)

var _ = time.Time{}

// Float64 is a type alias for float64
type Float64 = float64

// Float64Flag is a flag with type float64
type Float64Flag struct {
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

	Value       Float64
	Destination *Float64
}

// Apply populates the flag given the flag set and environment
func (f *Float64Flag) Apply(set *flag.FlagSet) error {
	return Apply(f, "float64", set)
}

// ApplyInputSourceValue applies a Float64 value to the flagSet if required
func (f *Float64Flag) ApplyInputSourceValue(context *Context, isc InputSourceContext) error {
	return ApplyInputSourceValue(f, context, isc)
}

// String returns a readable representation of this value
// (for usage defaults)
func (f *Float64Flag) String() string {
	return FlagStringer(f)
}

// Float64 looks up the value of a local Float64Flag, returns
// an empty value if not found
func (c *Context) Float64(name string) float64 {
	return c.Lookup(name, *new(Float64)).(float64)
}
