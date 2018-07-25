package cmd

import (
	"time"
	"sort"
	"github.com/urfave/cli"
	"os"
)

func RunCmd() {
	app := cli.NewApp()
	app.Compiled = time.Now()
	app.Authors = []cli.Author{{Name: "barry", Email: "kooksee@163.com"}}
	app.Flags = []cli.Flag{
		isDevflag(),
	}
	app.Commands = []cli.Command{
		DaemonCmd(),
		AccountCmd(),
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	if err := app.Run(os.Args); err != nil {
		logger.Error(err.Error())
	}
}
