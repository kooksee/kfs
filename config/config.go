package config

import (
	"os"
	log "github.com/inconshreveable/log15"
	"sync"
	"mu/cmn"
	"github.com/kooksee/kdb"
)

var once1 sync.Once

func (t *Config) GetDb() *kdb.KDB {
	if t.db == nil {
		panic("please init db")
	}
	return t.db
}

func GetCfg() *Config {
	if instance == nil {
		panic("please init config")
	}
	return instance
}

func homeDir(defaultHome string) string {
	if len(os.Args) > 2 && os.Args[len(os.Args)-2] == "--home" {
		defaultHome = os.Args[len(os.Args)-1]
		os.Args = os.Args[:len(os.Args)-2]
	}
	return defaultHome
}

func (t *Config) Log() log.Logger {
	if t.l == nil {
		panic("please init log")
	}
	return t.l
}

func NewCfg(defaultHomeDir string) *Config {
	defaultHomeDir = homeDir(defaultHomeDir)
	instance = &Config{}

	instance.home = defaultHomeDir
	instance.LogLevel = "debug"

	cmn.EnsureDir(instance.home, os.FileMode(0755))

	return instance
}
