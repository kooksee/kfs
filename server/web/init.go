package web

import (
	"github.com/json-iterator/go"
	"fmt"
	"github.com/kooksee/kfs/config"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
	cfg  *config.Config
)

func Init() {
	cfg = config.GetCfg()
}

func f(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

func str(s interface{}) string {
	return fmt.Sprintf("%s", s)
}

type m map[string]interface{}
