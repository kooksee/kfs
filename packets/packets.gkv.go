package sp2p

import "github.com/kooksee/kfs/sp2p"

/*
通过广播的方式进行数据存储,同时随机抽样检测数据的一致性
 */

var gkvPrefix = []byte("gkv")

type GKVSetReq struct{ kv }

func (t *GKVSetReq) T() byte        { return GKVSetReqT }
func (t *GKVSetReq) String() string { return GKVSetReqS }
func (t *GKVSetReq) OnHandle(p *sp2p.SP2p, msg *KMsg) {
	if err := GetDb().KHash(gkvPrefix).Set(t.K, t.V); err != nil {
		GetLog().Error(err.Error())
	}

	// 随机广播
	for _, node := range p.tab.FindRandomNodes(cfg.NodeBroadcastNumber) {
		p.writeTx(&KMsg{
			FAddr: msg.FAddr,
			Data:  msg.Data,
			TAddr: node.AddrString(),
		})
	}
}

type GKVGetReq struct{ kv }

func (t *GKVGetReq) T() byte        { return GKVGetReqT }
func (t *GKVGetReq) String() string { return GKVGetReqS }
func (t *GKVGetReq) OnHandle(p *SP2p, msg *KMsg) {
	v, _ := GetDb().KHash(gkvPrefix).Get(t.K)
	if v == nil || len(v) == 0 {
		for _, node := range p.tab.FindRandomNodes(8) {
			p.writeTx(&KMsg{
				Data:  msg.Data,
				FAddr: msg.FAddr,
				TAddr: node.AddrString(),
			})
		}
		return
	}

	resp := &KVGetResp{}
	resp.K = t.K
	resp.V = v
	p.Write(&KMsg{Data: resp, TAddr: msg.FAddr})
}

type GKVGetResp struct{ kv }

func (t *GKVGetResp) T() byte        { return GKVGetRespT }
func (t *GKVGetResp) String() string { return GKVGetRespS }
func (t *GKVGetResp) OnHandle(p *SP2p, msg *KMsg) {
	if err := GetDb().KHash(gkvPrefix).Set(t.K, t.V); err != nil {
		GetLog().Error(err.Error())
	}
}
