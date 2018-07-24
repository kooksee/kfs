package packets

import (
	"github.com/kooksee/kfs/sp2p"
	"github.com/inconshreveable/log15"
	"github.com/kooksee/kfs/config"
)

var (
	logger log15.Logger
	cfg    *config.Config
)

func Init() {
	cfg = config.GetCfg()
	logger = config.Log().New("package", "packets")

	sp2p.GetHManager().Registry(
		KVSetReq{},
		KVGetReq{},
		KVGetResp{},
	)
}
