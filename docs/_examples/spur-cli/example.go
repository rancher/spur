package main

import (
	"fmt"

	"github.com/rancher/spur/cli"
	"github.com/rancher/spur/cli/altsrc"
)

func main() {
	var (
		defaultValue = []string{"Hello", "world"}
		destination  = &[]string{"test", "ok"}
	)

	fmt.Printf("--- Start --------------------------------\n")
	fmt.Printf("Have test-string-slice value: %v\n", defaultValue)
	fmt.Printf("Have test-string-slice destination: %v\n", *destination)

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
		Before: cli.InitAllInputSource(altsrc.NewConfigFromFlag("load")),
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:        "test-string-slice",
				Aliases:     []string{"s"},
				EnvVars:     []string{"SLICE"},
				Value:       defaultValue,
				Destination: destination,
			},
			&cli.StringFlag{
				Name:  "load",
				Value: "conf.json",
			},
		},
	}).RunAndExitOnError()
}
