package config

import (
	"github.com/inconshreveable/log15"
	"os"
	"sync"
	"github.com/kooksee/kdb"
	"path/filepath"
)

var (
	once     sync.Once
	instance *Config
)

type Config struct {
	l        log15.Logger
	home     string
	db       *kdb.KDB
	IsDev    bool
	LogLevel string
}

func (t *Config) InitLog() {
	t.l = log15.New()
	ll, err := log15.LvlFromString(t.LogLevel)
	if err != nil {
		panic(err.Error())
	}
	t.l.SetHandler(log15.LvlFilterHandler(ll, log15.StreamHandler(os.Stdout, log15.TerminalFormat())))
}

func (t *Config) InitDb() {
	kdb.InitKdb(filepath.Join(t.home, ""))
	t.Db = kdb.GetKdb()
}

func (t *Config) MustNotErr(errs ... error) {
	for _, err := range errs {
		if err != nil {
			t.l.Error(err.Error())
			t.l.Error("看看看看看")
			panic(err.Error())
		}
	}
}
