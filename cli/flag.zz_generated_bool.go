package cli

import (
	"time"

	"github.com/rancher/spur/flag"
)

var _ = time.Time{}

// Bool is a type alias for bool
type Bool = bool

// BoolFlag is a flag with type bool
type BoolFlag struct {
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

	Value       Bool
	Destination *Bool
}

// Apply populates the flag given the flag set and environment
func (f *BoolFlag) Apply(set *flag.FlagSet) error {
	return Apply(f, "bool", set)
}

// ApplyInputSourceValue applies a Bool value to the flagSet if required
func (f *BoolFlag) ApplyInputSourceValue(context *Context, isc InputSourceContext) error {
	return ApplyInputSourceValue(f, context, isc)
}

// String returns a readable representation of this value
// (for usage defaults)
func (f *BoolFlag) String() string {
	return FlagStringer(f)
}

// Bool looks up the value of a local BoolFlag, returns
// an empty value if not found
func (c *Context) Bool(name string) bool {
	return c.Lookup(name, *new(Bool)).(bool)
}
