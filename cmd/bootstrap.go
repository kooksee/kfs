package cmd

import (
	"github.com/urfave/cli"
)

func BootstrapRmCmd() cli.Command {
	return cli.Command{
		Name:    "rm",
		Usage:   "rm peers",
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			return nil
		},
	}
}

func BootstrapLsCmd() cli.Command {
	return cli.Command{
		Name:    "ls",
		Usage:   "ls peers",
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			return nil
		},
	}
}

func BootstrapAddCmd() cli.Command {
	return cli.Command{
		Name:    "add",
		Usage:   "add peer",
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			return nil
		},
	}
}

func BootstrapCmd() cli.Command {
	return cli.Command{
		Name:    "bootstrap",
		Aliases: []string{"bs"},
		Usage:   "manage peers",
		Flags:   []cli.Flag{},
		Subcommands: []cli.Command{
			BootstrapAddCmd(),
			BootstrapLsCmd(),
			BootstrapRmCmd(),
		},
		Action: func(c *cli.Context) error {
			return nil
		},
	}
}
