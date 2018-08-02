package core

import (
	"github.com/inconshreveable/log15"
	"github.com/kooksee/kdb"
	"github.com/kooksee/kfs/config"
)

var (
	logger log15.Logger
	cfg    *config.Config
	kvDb   kdb.IKHash
	metaDb kdb.IKHash

	kvPrefix   = []byte("kv")
	metaPrefix = []byte("meta")
)

func Init() {
	cfg = config.GetCfg()
	logger = cfg.Log().New("package", "api.core")
	kvDb = cfg.GetDb().KHash(kvPrefix)
	metaDb = cfg.GetDb().KHash(metaPrefix)
}
