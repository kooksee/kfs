package sp2p

func (s *SP2p) Write(msg *KMsg) {
	go s.writeTx(msg)
}

func (s *SP2p) GetTable() *Table {
	return s.tab
}

func (s *SP2p) GetNode() string {
	return s.tab.selfNode.String()
}

func (s *SP2p) GetNodes() []string {
	return s.tab.GetRawNodes()
}

func (s *SP2p) TableSize() int {
	return s.tab.Size()
}

func (s *SP2p) UpdateNode(rawUrl string) error {
	n, err := ParseNode(rawUrl)
	if err != nil {
		return err
	}
	s.tab.UpdateNode(n)
	return nil
}
func (s *SP2p) DeleteNode(id string) error {
	n, err := HexToHash(id)
	if err != nil {
		return err
	}
	s.tab.DeleteNode(n)
	return nil
}

func (s *SP2p) AddNode(rawUrl string) error {
	n, err := ParseNode(rawUrl)
	if err != nil {
		return err
	}
	s.tab.AddNode(n)
	return nil
}

func (s *SP2p) FindMinDisNodes(targetID string, n int) (nodes []string, err error) {
	h, err := HexToHash(targetID)
	if err != nil {
		return nil, err
	}

	for _, n := range s.tab.FindMinDisNodes(h, n) {
		nodes = append(nodes, n.String())
	}
	return nodes, nil
}

func (s *SP2p) FindRandomNodes(n int) (nodes []string) {
	for _, n := range s.tab.FindRandomNodes(n) {
		nodes = append(nodes, n.String())
	}
	return
}

func (s *SP2p) FindNodeWithTargetBySelf(d string) (nodes []string) {
	for _, n := range s.tab.FindNodeWithTargetBySelf(StringToHash(d)) {
		nodes = append(nodes, n.String())
	}
	return
}

func (s *SP2p) FindNodeWithTarget(targetId string, measure string) (nodes []string) {
	for _, n := range s.tab.FindNodeWithTarget(StringToHash(targetId), StringToHash(measure)) {
		nodes = append(nodes, n.String())
	}
	return
}

func (s *SP2p) PingN() {
	go s.pingN()
}

func (s *SP2p) PingNode(taddr, tid string) {
	go s.pingNode(taddr, tid)
}

func (s *SP2p) FindN() {
	go s.findN()
}

func (s *SP2p) FindNode(taddr, tid string, n int) {
	go s.findNode(taddr, tid, n)
}
