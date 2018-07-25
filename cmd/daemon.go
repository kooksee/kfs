package cmd

import (
	"github.com/urfave/cli"
	"github.com/kooksee/kfs/sp2p"
)

func DaemonCmd() cli.Command {
	return cli.Command{
		Name:    "daemon",
		Aliases: []string{"d"},
		Usage:   "start kfs daemon",
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			f := sp2p.InitCfg().InitLog(logger).InitDb(cfg.GetDb())
			f.PriV = nil

			sp2p.NewSP2p()

			return nil
		},
	}
}
