package cmd

import (
	"github.com/inconshreveable/log15"
	"github.com/urfave/cli"
	"github.com/kooksee/kfs/config"
)

var (
	logger log15.Logger
	cfg    *config.Config
)

func Init() {
	cfg = config.GetCfg()
	logger = config.Log().New("package", "cmd")
}

var serverPort int

//func accountNumFlag() cli.IntFlag { return cli.IntFlag{Name: "num", Value: accountNum, Destination: &accountNum, Usage: "生成账号数"} }
func aerverAddr() cli.IntFlag { return cli.IntFlag{Name: "port", Value: serverPort, Destination: &serverPort, Usage: "端口"} }
func isDevflag() cli.BoolFlag { return cli.BoolFlag{Name: "debug", Destination: &cfg.IsDev, Usage: "debug mode"} }
