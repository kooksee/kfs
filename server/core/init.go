package core

import (
	"github.com/inconshreveable/log15"
	"github.com/kooksee/kdb"
	"github.com/kooksee/kfs/config"
)

var (
	logger log15.Logger
	cfg    *config.Config
	kvDb   *kdb.KHash
	metaDb *kdb.KHash

	kvPrefix   = []byte("kv")
	metaPrefix = []byte("meta")
)

func Init() {
	cfg = config.GetCfg()
	logger = cfg.Log().New("package", "packets")
	kvDb = cfg.GetDb().KHash(kvPrefix)
	metaDb = cfg.GetDb().KHash(metaPrefix)
}