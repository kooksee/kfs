package sp2p

type FindNodeReq struct {
	N int `json:"n,omitempty"`
}

func (t *FindNodeReq) T() byte        { return FindNodeReqT }
func (t *FindNodeReq) String() string { return FindNodeReqS }
func (t *FindNodeReq) OnHandle(p *SP2p, msg *KMsg) {

	node, err := NodeFromKMsg(msg)
	if err != nil {
		GetLog().Error("NodeFromKMsg error", "err", err)
		return
	}
	go p.tab.UpdateNode(node)

	ns := make([]string, 0)

	// 最多不能超过16
	if t.N > 16 {
		t.N = 16
	}

	for _, n := range p.tab.FindMinDisNodes(node.ID, t.N) {
		ns = append(ns, n.String())
	}
	p.Write(&KMsg{TAddr: msg.Addr, Data: &FindNodeResp{Nodes: ns}})
}

type FindNodeResp struct {
	Nodes []string `json:"nodes,omitempty"`
}

func (t *FindNodeResp) T() byte        { return FindNodeRespT }
func (t *FindNodeResp) String() string { return FindNodeRespS }
func (t *FindNodeResp) OnHandle(p *SP2p, msg *KMsg) {
	for _, n := range t.Nodes {
		node, err := ParseNode(n)
		if err != nil {
			GetLog().Error("parse node error", "err", err)
			continue
		}
		p.tab.UpdateNode(node)
	}
}
