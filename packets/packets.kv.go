package packets

import (
	"github.com/kooksee/sp2p"
)

/*
采用分片的方式进行kv存储,同时定时抽样的方式检测自己的数据是否有合适的节点可以存储
 */

type kv struct {
	K []byte `json:"k,omitempty"`
	V []byte `json:"v,omitempty"`
}

type KVSetReq struct{ kv }

func (t *KVSetReq) T() byte        { return KVSetReqT }
func (t *KVSetReq) String() string { return KVSetReqS }
func (t *KVSetReq) OnHandle(p sp2p.ISP2P, msg *sp2p.KMsg) {
	nodes := p.FindNodeWithTargetBySelf(sp2p.BytesToHash(t.K).Hex())

	if len(nodes) < 3 {
		if err := kvDb.Set(t.K, t.V); err != nil {
			logger.Error("kvset error", "err", err)
		}
		return
	}

	for _, node := range nodes {
		p.Write(&sp2p.KMsg{FAddr: msg.FAddr, Data: msg.Data, TAddr: node})
	}
}

type KVGetReq struct{ kv }

func (t *KVGetReq) T() byte        { return KVGetReqT }
func (t *KVGetReq) String() string { return KVGetReqS }
func (t *KVGetReq) OnHandle(p sp2p.ISP2P, msg *sp2p.KMsg) {
	nodes := p.FindNodeWithTargetBySelf(sp2p.BytesToHash(t.K))
	if len(nodes) < 3 {
		resp := &KVGetResp{}
		resp.K = t.K
		resp.V, _ = kvDb.Get(t.K)

		if len(resp.V) != 0 {
			p.Write(&sp2p.KMsg{Data: resp, TAddr: msg.Addr, TID: msg.ID})
			return
		}
	}

	for _, node := range nodes {
		p.Write(&sp2p.KMsg{Data: msg.Data, Addr: msg.Addr, ID: msg.ID, TID: node.ID.ToHex(), TAddr: node.AddrString()})
	}
}

type KVGetResp struct{ kv }

func (t *KVGetResp) T() byte        { return KVGetRespT }
func (t *KVGetResp) String() string { return KVGetRespS }
func (t *KVGetResp) OnHandle(p sp2p.ISP2P, msg *sp2p.KMsg) {
	if err := kvDb.Set(t.K, t.V); err != nil {
		logger.Error(err.Error())
	}
}
