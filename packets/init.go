package packets

import (
	"github.com/kooksee/sp2p"
	"github.com/inconshreveable/log15"
	"github.com/kooksee/kfs/config"
	"github.com/kooksee/kdb"
)

var (
	logger log15.Logger
	cfg    *config.Config
	kvDb   kdb.IKHash

	kvPrefix = []byte("kv")
)

func Init() {
	cfg = config.GetCfg()
	logger = cfg.Log().New("package", "packets")
	kvDb = cfg.GetDb().KHash(kvPrefix)

	sp2p.RegistryHandlers(
		KVSetReq{},
		KVGetReq{},
		KVGetResp{},

		ChatReq{},
		ChatResp{},
	)
}
