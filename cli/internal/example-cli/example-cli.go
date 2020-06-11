// minimal example CLI used for binary size checking

package main

import (
	"github.com/rancher/spur/cli"
)

func main() {
	(&cli.App{}).Run([]string{"--help"})
}
