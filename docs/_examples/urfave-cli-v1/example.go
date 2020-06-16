package main

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/urfave/cli/altsrc"
)

func main() {
	var (
		defaultValue cli.StringSlice = []string{"Hello", "world"}
	)

	fmt.Printf("--- Start --------------------------------\n")
	fmt.Printf("Have test-string-slice value: %v\n", defaultValue)

	flags := []cli.Flag{
		altsrc.NewStringSliceFlag(cli.StringSliceFlag{
			Name:   "test-string-slice,s",
			EnvVar: "SLICE",
			Value:  &defaultValue,
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:  "load",
			Value: "conf.json",
		}),
	}
	(&cli.App{
		EnableBashCompletion: true,
		Action: func(c *cli.Context) error {
			fmt.Printf("--- Action --------------------------------\n")
			fmt.Printf("Have test-string-slice value: %v\n", defaultValue)
			fmt.Printf("Have test-string-slice context: %v\n", c.StringSlice("test-string-slice"))
			fmt.Printf("Have load context: %v\n", c.String("load"))
			return nil
		},
		Before: altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("load")),
		Flags:  flags,
	}).RunAndExitOnError()
}
