package sp2p

func init() {
	GetHManager().Registry(
		KVSetReq{},
		KVGetReq{},
		KVGetResp{},

		GKVSetReq{},
		GKVGetReq{},
		GKVGetResp{},
	)
}
