package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func main() {
	var (
		defaultValue *cli.StringSlice = cli.NewStringSlice("Hello", "world")
		destination  *cli.StringSlice = cli.NewStringSlice("test", "ok")
	)

	fmt.Printf("--- Start --------------------------------\n")
	fmt.Printf("Have test-string-slice value: %v\n", defaultValue)
	fmt.Printf("Have test-string-slice destination: %v\n", *destination)

	flags := []cli.Flag{
		altsrc.NewStringSliceFlag(&cli.StringSliceFlag{
			Name:        "test-string-slice",
			Aliases:     []string{"s"},
			EnvVars:     []string{"SLICE"},
			Value:       defaultValue,
			Destination: destination,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  "load",
			Value: "conf.json",
		}),
	}
	(&cli.App{
		EnableBashCompletion: true,
		Action: func(c *cli.Context) error {
			fmt.Printf("--- Action --------------------------------\n")
			fmt.Printf("Have test-string-slice value: %v\n", defaultValue)
			fmt.Printf("Have test-string-slice destination: %v\n", *destination)
			fmt.Printf("Have test-string-slice context: %v\n", c.StringSlice("test-string-slice"))
			fmt.Printf("Have load context: %v\n", c.String("load"))
			return nil
		},
		Before: altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("load")),
		Flags:  flags,
	}).RunAndExitOnError()
}
