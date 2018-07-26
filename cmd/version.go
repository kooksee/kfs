package cmd

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/kooksee/kfs/version"
)

func VersionCmd() cli.Command {
	return cli.Command{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Show version info",
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			fmt.Println("kfs version", version.Version)
			fmt.Println("kfs commit version", version.GitCommit)
			fmt.Println("kfs build version", version.BuildVersion)
			return nil
		},
	}
}
