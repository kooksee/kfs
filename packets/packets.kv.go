package packets

import "github.com/kooksee/kfs/sp2p"

/*
采用分片的方式进行kv存储,同时定时抽样的方式检测自己的数据是否有合适的节点可以存储
 */

var kvPrefix = []byte("kv")

type kv struct {
	K       []byte `json:"k,omitempty"`
	V       []byte `json:"v,omitempty"`
	Expired int    `json:"expired,omitempty"`
	Time    int    `json:"time,omitempty"`
}

type KVSetReq struct{ kv }

func (t *KVSetReq) T() byte        { return KVSetReqT }
func (t *KVSetReq) String() string { return KVSetReqS }
func (t *KVSetReq) OnHandle(p *sp2p.SP2p, msg *sp2p.KMsg) {
	nodes := p.GetTable().FindNodeWithTargetBySelf(sp2p.BytesToHash(t.K))
	if len(nodes) < cfg.NodePartitionNumber {
		if err := cfg.GetDb().KHash(kvPrefix).Set(t.K, t.V); err != nil {
			GetLog().Error("kvset error", "err", err)
		}
		return
	}

	for _, node := range nodes {
		p.writeTx(&KMsg{FAddr: msg.FAddr, Data: msg.Data, TAddr: node.AddrString()})
	}
}

type KVGetReq struct{ kv }

func (t *KVGetReq) T() byte        { return KVGetReqT }
func (t *KVGetReq) String() string { return KVGetReqS }
func (t *KVGetReq) OnHandle(p *SP2p, msg *KMsg) {
	nodes := p.GetTable().FindNodeWithTargetBySelf(BytesToHash(t.K))
	if len(nodes) < cfg.NodePartitionNumber {
		resp := &KVGetResp{}
		resp.K = t.K
		resp.V, _ = GetDb().KHash(kvPrefix).Get(t.K)

		p.writeTx(&KMsg{Data: resp, TAddr: msg.FAddr})
		return
	}

	for _, node := range nodes {
		p.writeTx(&KMsg{Data: msg.Data, FAddr: msg.FAddr, TAddr: node.AddrString()})
	}
}

type KVGetResp struct{ kv }

func (t *KVGetResp) T() byte        { return KVGetRespT }
func (t *KVGetResp) String() string { return KVGetRespS }
func (t *KVGetResp) OnHandle(p *SP2p, msg *KMsg) {
	if err := GetDb().KHash(kvPrefix).Set(t.K, t.V); err != nil {
		GetLog().Error(err.Error())
	}
}
