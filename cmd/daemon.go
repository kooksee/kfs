package cmd

import (
	"github.com/urfave/cli"
	"github.com/kooksee/kfs/sp2p"
	"github.com/kooksee/kfs/packets"
)

func DaemonCmd() cli.Command {
	return cli.Command{
		Name:    "daemon",
		Aliases: []string{"d"},
		Usage:   "start kfs daemon",
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			f := sp2p.InitCfg().InitLog(logger).InitDb(cfg.GetDb())
			f.Seeds = cfg.Seeds
			f.Adds = cfg.Adds
			f.KeyStore = cfg.GetKeyStore()

			// 注册handle
			packets.Init()

			// 启动p2p通讯
			sp2p.NewSP2p()

			return nil
		},
	}
}
