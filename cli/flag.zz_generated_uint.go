package cli

import (
	"time"

	"github.com/rancher/spur/flag"
)

var _ = time.Time{}

// Uint is a type alias for uint
type Uint = uint

// UintFlag is a flag with type uint
type UintFlag struct {
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

	Value       Uint
	Destination *Uint
}

// Apply populates the flag given the flag set and environment
func (f *UintFlag) Apply(set *flag.FlagSet) error {
	return Apply(f, "uint", set)
}

// ApplyInputSourceValue applies a Uint value to the flagSet if required
func (f *UintFlag) ApplyInputSourceValue(context *Context, isc InputSourceContext) error {
	return ApplyInputSourceValue(f, context, isc)
}

// String returns a readable representation of this value
// (for usage defaults)
func (f *UintFlag) String() string {
	return FlagStringer(f)
}

// Uint looks up the value of a local UintFlag, returns
// an empty value if not found
func (c *Context) Uint(name string) uint {
	return c.Lookup(name, *new(Uint)).(uint)
}
