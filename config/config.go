package config

import (
	"os"
	log "github.com/inconshreveable/log15"
	"sync"
	"mu/cmn"
)

var once1 sync.Once

func GetCfg() *Config {
	if instance == nil {
		panic("please init config")
	}
	return instance
}

func GetHomeDir(defaultHome string) string {
	if len(os.Args) > 2 && os.Args[len(os.Args)-2] == "--home" {
		defaultHome = os.Args[len(os.Args)-1]
		os.Args = os.Args[:len(os.Args)-2]
	}
	return defaultHome
}

func Log() log.Logger {
	cfg := GetCfg()
	if cfg.l == nil {
		panic("please init log")
	}
	return cfg.l
}

func NewCfg(defaultHomeDir string) *Config {
	defaultHomeDir = GetHomeDir(defaultHomeDir)
	instance = &Config{}

	instance.home = defaultHomeDir
	instance.LogLevel = "debug"

	cmn.EnsureDir(instance.home, os.FileMode(0755))

	return instance
}
