package main

import (
	"github.com/kooksee/kfs/config"
	"github.com/kooksee/kfs/cmd"
)

func main() {
	cfg := config.NewCfg("kdata")
	cfg.InitLog()
	cfg.InitDb()

	cmd.Init()
	cmd.RunCmd()
}
