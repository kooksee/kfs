package sp2p

func init() {
	GetHManager().Registry(
		PingReq{},
		FindNodeReq{},
		FindNodeResp{},
	)
}
