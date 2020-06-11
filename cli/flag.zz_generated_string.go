package cli

import (
	"time"

	"github.com/rancher/spur/flag"
)

var _ = time.Time{}

// String is a type alias for string
type String = string

// StringFlag is a flag with type string
type StringFlag struct {
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

	Value       String
	Destination *String
}

// Apply populates the flag given the flag set and environment
func (f *StringFlag) Apply(set *flag.FlagSet) error {
	return Apply(f, "string", set)
}

// ApplyInputSourceValue applies a String value to the flagSet if required
func (f *StringFlag) ApplyInputSourceValue(context *Context, isc InputSourceContext) error {
	return ApplyInputSourceValue(f, context, isc)
}

// String returns a readable representation of this value
// (for usage defaults)
func (f *StringFlag) String() string {
	return FlagStringer(f)
}

// String looks up the value of a local StringFlag, returns
// an empty value if not found
func (c *Context) String(name string) string {
	return c.Lookup(name, *new(String)).(string)
}
