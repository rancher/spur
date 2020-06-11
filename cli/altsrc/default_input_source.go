package altsrc

import "github.com/rancher/spur/cli"

// defaultInputSource creates a default cli.InputSourceContext.
func defaultInputSource() (cli.InputSourceContext, error) {
	return &MapInputSource{file: "", valueMap: map[interface{}]interface{}{}}, nil
}
