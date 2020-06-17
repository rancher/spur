package altsrc

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/rancher/spur/cli"
	"github.com/rancher/spur/flag"
)

func TestConfigFileTest(t *testing.T) {
	app := &cli.App{}
	set := flag.NewFlagSet("test", 0)
	ioutil.WriteFile("current.yaml", []byte("test: 15"), 0666)
	defer os.Remove("current.yaml")
	test := []string{"test-cmd", "--load", "current.yaml"}
	set.Parse(test)

	c := cli.NewContext(app, set, nil)

	command := &cli.Command{
		Name:        "test-cmd",
		Aliases:     []string{"tc"},
		Usage:       "this is for testing",
		Description: "testing",
		Action: func(c *cli.Context) error {
			val := c.Int("test")
			expect(t, val, 15)
			return nil
		},
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "test"},
			&cli.StringFlag{Name: "load"}},
	}
	command.Before = cli.InitAllInputSource(NewConfigFromFlag("load"))
	err := command.Run(c)

	expect(t, err, nil)
}

func TestConfigNoFileTest(t *testing.T) {
	app := &cli.App{}
	set := flag.NewFlagSet("test", 0)
	test := []string{"test-cmd", "--load", ".doesNotExist.yaml"}
	set.Parse(test)

	c := cli.NewContext(app, set, nil)

	command := &cli.Command{
		Name:        "test-cmd",
		Aliases:     []string{"tc"},
		Usage:       "this is for testing",
		Description: "testing",
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "test"},
			&cli.StringFlag{Name: "load"}},
	}
	command.Before = cli.InitAllInputSource(NewConfigFromFlag("load"))
	err := command.Run(c)

	refute(t, err, nil)
}

func TestConfigDefaultFileTest(t *testing.T) {
	app := &cli.App{}
	set := flag.NewFlagSet("test", 0)
	ioutil.WriteFile("current.yaml", []byte("test: 15"), 0666)
	defer os.Remove("current.yaml")
	test := []string{"test-cmd"}
	set.Parse(test)

	c := cli.NewContext(app, set, nil)

	command := &cli.Command{
		Name:        "test-cmd",
		Aliases:     []string{"tc"},
		Usage:       "this is for testing",
		Description: "testing",
		Action: func(c *cli.Context) error {
			val := c.Int("test")
			expect(t, val, 15)
			return nil
		},
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "test"},
			&cli.StringFlag{
				Name:  "load",
				Value: "current.yaml",
			}},
	}
	command.Before = cli.InitAllInputSource(NewConfigFromFlag("load"))
	err := command.Run(c)

	expect(t, err, nil)
}

func TestConfigDefaultInvalidFileTest(t *testing.T) {
	app := &cli.App{}
	set := flag.NewFlagSet("test", 0)
	ioutil.WriteFile("current.yaml", []byte("test: [15"), 0666)
	defer os.Remove("current.yaml")
	test := []string{"test-cmd"}
	set.Parse(test)

	c := cli.NewContext(app, set, nil)

	command := &cli.Command{
		Name:        "test-cmd",
		Aliases:     []string{"tc"},
		Usage:       "this is for testing",
		Description: "testing",
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "test"},
			&cli.StringFlag{
				Name:  "load",
				Value: "current.yaml",
			}},
	}
	command.Before = cli.InitAllInputSource(NewConfigFromFlag("load"))
	err := command.Run(c)

	refute(t, err, nil)
}

func TestConfigFileTestGlobalEnvVarWins(t *testing.T) {
	app := &cli.App{}
	set := flag.NewFlagSet("test", 0)
	ioutil.WriteFile("current.yaml", []byte("test: 15"), 0666)
	defer os.Remove("current.yaml")

	os.Setenv("THE_TEST", "10")
	defer os.Setenv("THE_TEST", "")
	test := []string{"test-cmd", "--load", "current.yaml"}
	set.Parse(test)

	c := cli.NewContext(app, set, nil)

	command := &cli.Command{
		Name:        "test-cmd",
		Aliases:     []string{"tc"},
		Usage:       "this is for testing",
		Description: "testing",
		Action: func(c *cli.Context) error {
			val := c.Int("test")
			expect(t, val, 10)
			return nil
		},
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "test", EnvVars: []string{"THE_TEST"}},
			&cli.StringFlag{Name: "load"}},
	}
	command.Before = cli.InitAllInputSource(NewConfigFromFlag("load"))

	err := command.Run(c)

	expect(t, err, nil)
}

func TestConfigFileTestGlobalEnvVarWinsNested(t *testing.T) {
	app := &cli.App{}
	set := flag.NewFlagSet("test", 0)
	ioutil.WriteFile("current.yaml", []byte(`top:
  test: 15`), 0666)
	defer os.Remove("current.yaml")

	os.Setenv("THE_TEST", "10")
	defer os.Setenv("THE_TEST", "")
	test := []string{"test-cmd", "--load", "current.yaml"}
	set.Parse(test)

	c := cli.NewContext(app, set, nil)

	command := &cli.Command{
		Name:        "test-cmd",
		Aliases:     []string{"tc"},
		Usage:       "this is for testing",
		Description: "testing",
		Action: func(c *cli.Context) error {
			val := c.Int("top.test")
			expect(t, val, 10)
			return nil
		},
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "top.test", EnvVars: []string{"THE_TEST"}},
			&cli.StringFlag{Name: "load"}},
	}
	command.Before = cli.InitAllInputSource(NewConfigFromFlag("load"))

	err := command.Run(c)

	expect(t, err, nil)
}

func TestConfigFileTestSpecifiedFlagWins(t *testing.T) {
	app := &cli.App{}
	set := flag.NewFlagSet("test", 0)
	ioutil.WriteFile("current.yaml", []byte("test: 15"), 0666)
	defer os.Remove("current.yaml")

	test := []string{"test-cmd", "--load", "current.yaml", "--test", "7"}
	set.Parse(test)

	c := cli.NewContext(app, set, nil)

	command := &cli.Command{
		Name:        "test-cmd",
		Aliases:     []string{"tc"},
		Usage:       "this is for testing",
		Description: "testing",
		Action: func(c *cli.Context) error {
			val := c.Int("test")
			expect(t, val, 7)
			return nil
		},
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "test"},
			&cli.StringFlag{Name: "load"}},
	}
	command.Before = cli.InitAllInputSource(NewConfigFromFlag("load"))

	err := command.Run(c)

	expect(t, err, nil)
}

func TestConfigFileTestSpecifiedFlagWinsNested(t *testing.T) {
	app := &cli.App{}
	set := flag.NewFlagSet("test", 0)
	ioutil.WriteFile("current.yaml", []byte(`top:
  test: 15`), 0666)
	defer os.Remove("current.yaml")

	test := []string{"test-cmd", "--load", "current.yaml", "--top.test", "7"}
	set.Parse(test)

	c := cli.NewContext(app, set, nil)

	command := &cli.Command{
		Name:        "test-cmd",
		Aliases:     []string{"tc"},
		Usage:       "this is for testing",
		Description: "testing",
		Action: func(c *cli.Context) error {
			val := c.Int("top.test")
			expect(t, val, 7)
			return nil
		},
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "top.test"},
			&cli.StringFlag{Name: "load"}},
	}
	command.Before = cli.InitAllInputSource(NewConfigFromFlag("load"))

	err := command.Run(c)

	expect(t, err, nil)
}

func TestConfigFileTestDefaultValueFileWins(t *testing.T) {
	app := &cli.App{}
	set := flag.NewFlagSet("test", 0)
	ioutil.WriteFile("current.yaml", []byte("test: 15"), 0666)
	defer os.Remove("current.yaml")

	test := []string{"test-cmd", "--load", "current.yaml"}
	set.Parse(test)

	c := cli.NewContext(app, set, nil)

	command := &cli.Command{
		Name:        "test-cmd",
		Aliases:     []string{"tc"},
		Usage:       "this is for testing",
		Description: "testing",
		Action: func(c *cli.Context) error {
			val := c.Int("test")
			expect(t, val, 15)
			return nil
		},
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "test", Value: 7},
			&cli.StringFlag{Name: "load"}},
	}
	command.Before = cli.InitAllInputSource(NewConfigFromFlag("load"))

	err := command.Run(c)

	expect(t, err, nil)
}

func TestConfigFileTestDefaultValueFileWinsNested(t *testing.T) {
	app := &cli.App{}
	set := flag.NewFlagSet("test", 0)
	ioutil.WriteFile("current.yaml", []byte(`top:
  test: 15`), 0666)
	defer os.Remove("current.yaml")

	test := []string{"test-cmd", "--load", "current.yaml"}
	set.Parse(test)

	c := cli.NewContext(app, set, nil)

	command := &cli.Command{
		Name:        "test-cmd",
		Aliases:     []string{"tc"},
		Usage:       "this is for testing",
		Description: "testing",
		Action: func(c *cli.Context) error {
			val := c.Int("top.test")
			expect(t, val, 15)
			return nil
		},
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "top.test", Value: 7},
			&cli.StringFlag{Name: "load"}},
	}
	command.Before = cli.InitAllInputSource(NewConfigFromFlag("load"))

	err := command.Run(c)

	expect(t, err, nil)
}

func TestConfigFileFlagHasDefaultGlobalEnvSetGlobalEnvWins(t *testing.T) {
	app := &cli.App{}
	set := flag.NewFlagSet("test", 0)
	ioutil.WriteFile("current.yaml", []byte("test: 15"), 0666)
	defer os.Remove("current.yaml")

	os.Setenv("THE_TEST", "11")
	defer os.Setenv("THE_TEST", "")

	test := []string{"test-cmd", "--load", "current.yaml"}
	set.Parse(test)

	c := cli.NewContext(app, set, nil)

	command := &cli.Command{
		Name:        "test-cmd",
		Aliases:     []string{"tc"},
		Usage:       "this is for testing",
		Description: "testing",
		Action: func(c *cli.Context) error {
			val := c.Int("test")
			expect(t, val, 11)
			return nil
		},
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "test", Value: 7, EnvVars: []string{"THE_TEST"}},
			&cli.StringFlag{Name: "load"}},
	}
	command.Before = cli.InitAllInputSource(NewConfigFromFlag("load"))
	err := command.Run(c)

	expect(t, err, nil)
}

func TestConfigFileFlagHasDefaultGlobalEnvSetGlobalEnvWinsNested(t *testing.T) {
	app := &cli.App{}
	set := flag.NewFlagSet("test", 0)
	ioutil.WriteFile("current.yaml", []byte(`top:
  test: 15`), 0666)
	defer os.Remove("current.yaml")

	os.Setenv("THE_TEST", "11")
	defer os.Setenv("THE_TEST", "")

	test := []string{"test-cmd", "--load", "current.yaml"}
	set.Parse(test)

	c := cli.NewContext(app, set, nil)

	command := &cli.Command{
		Name:        "test-cmd",
		Aliases:     []string{"tc"},
		Usage:       "this is for testing",
		Description: "testing",
		Action: func(c *cli.Context) error {
			val := c.Int("top.test")
			expect(t, val, 11)
			return nil
		},
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "top.test", Value: 7, EnvVars: []string{"THE_TEST"}},
			&cli.StringFlag{Name: "load"}},
	}
	command.Before = cli.InitAllInputSource(NewConfigFromFlag("load"))
	err := command.Run(c)

	expect(t, err, nil)
}
