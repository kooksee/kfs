package packets

import (
	"github.com/kooksee/sp2p"
)

/*
采用分片的方式进行kv存储,同时定时抽样的方式检测自己的数据是否有合适的节点可以存储
 */

type kv struct {
	sp2p.IMessage

	K []byte `json:"k,omitempty"`
	V []byte `json:"v,omitempty"`
}

type kVSetReq struct{ kv }

func (t *kVSetReq) T() byte        { return kVSetReqT }
func (t *kVSetReq) String() string { return kVSetReqS }
func (t *kVSetReq) OnHandle(p sp2p.ISP2P, msg *sp2p.KMsg) error {
	nodes := p.FindNodeWithTargetBySelf(sp2p.BytesToHash(t.K).Hex())

	if len(nodes) < 3 {
		return kvDb.Set(t.K, t.V)
	}

	for _, node := range nodes {
		p.Write(&sp2p.KMsg{FN: msg.FN, Data: msg.Data, TN: node})
	}

	return nil
}

type kVGetReq struct{ kv }

func (t *kVGetReq) T() byte        { return kVGetReqT }
func (t *kVGetReq) String() string { return kVGetReqS }
func (t *kVGetReq) OnHandle(p sp2p.ISP2P, msg *sp2p.KMsg) error {
	nodes := p.FindNodeWithTargetBySelf(sp2p.BytesToHash(t.K).Hex())
	if len(nodes) < 3 {
		resp := &kVGetResp{}
		resp.K = t.K
		resp.V, _ = kvDb.Get(t.K)

		if len(resp.V) != 0 {
			p.Write(&sp2p.KMsg{Data: resp, TN: msg.TN})
			return nil
		}
	}

	for _, node := range nodes {
		p.Write(&sp2p.KMsg{Data: msg.Data, FN: msg.FN, TN: node})
	}

	return nil
}

type kVGetResp struct{ kv }

func (t *kVGetResp) T() byte        { return kVGetRespT }
func (t *kVGetResp) String() string { return kVGetRespS }
func (t *kVGetResp) OnHandle(p sp2p.ISP2P, msg *sp2p.KMsg) error {
	return kvDb.Set(t.K, t.V)
}
