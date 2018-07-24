package sp2p

import (
	"net"
	"github.com/satori/go.uuid"
	"github.com/ethereum/go-ethereum/crypto"
	"io"
	"time"
	"strings"
)

func NewSP2p() *SP2p {
	logger := GetLog()

	taddr, err := net.ResolveTCPAddr("tcp", cfg.Adds[0])
	if err != nil {
		panic(err.Error())
	}

	p2p := &SP2p{
		txRC:      make(chan *KMsg, 10000),
		txWC:      make(chan *KMsg, 10000),
		localAddr: taddr,
		laddr:     cfg.Adds[0],
	}

	uad, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	if err != nil {
		panic(err.Error())
	}
	cnn, err := net.ListenUDP("udp", uad)
	if err != nil {
		panic(err.Error())
	}
	p2p.rconn = cnn

	// 连接服务端等待
	// todo: 同时连接多个服务端切换
	// todo: 尝试连接不同的服务端
	for {
		if conn, err := net.DialTCP("tcp", nil, p2p.localAddr); err != nil {
			logger.Error(Errs(Fmt("udp %s listen error", p2p.localAddr), err.Error()))
			time.Sleep(time.Second * 2)
		} else {
			p2p.conn = conn
			break
		}
	}

	// 生成node id
	nodeId := MustBytesID(crypto.FromECDSAPub(&cfg.priv.PublicKey))
	p2p.nid = nodeId.ToHex()

	logger.Debug("node id", "id", nodeId)
	logger.Debug("create table", "table")

	p2p.tab = newTable(nodeId, p2p.localAddr)

	// 把seeds添加到集群中
	for _, s := range cfg.Seeds {
		p2p.tab.UpdateNode(MustParseNode(s))
	}

	go p2p.accept()
	go p2p.loop()
	go p2p.genUUID()

	return p2p
}

type SP2p struct {
	tab       *Table
	txRC      chan *KMsg
	txWC      chan *KMsg
	conn      *net.TCPConn
	rconn     *net.UDPConn
	localAddr *net.TCPAddr
	laddr     string
	nid       string
}

// 生成uuid的队列
func (s *SP2p) genUUID() {
	for {
		uid, err := uuid.NewV4()
		if err == nil {
			cfg.uuidC <- uid.String()
		}
	}
}

func (s *SP2p) GetAddr() string {
	return s.laddr
}

func (s *SP2p) pingRaley() {
	m := &KMsg{
		TID:   s.tab.selfNode.ID.ToHex(),
		TAddr: s.GetAddr(),
	}
	s.conn.Write(m.Dumps())
}

func (s *SP2p) loop() {
	for {
		select {
		case <-cfg.RelayTick.C:
			// 定时访问一下中继
			go s.pingRaley()
		case <-cfg.FindNodeTick.C:
			// 定时查找其他的节点
			go s.findN()
		case <-cfg.PingTick.C:
			// 定时ping其他的节点
			go s.pingN()
		case <-cfg.NtpTick.C:
			// 定时检查本地时间同步
			go checkClockDrift()
		case tx := <-s.txRC:
			// 处理tx
			go tx.Data.OnHandle(s, tx)
		case tx := <-s.txWC:
			// 处理发送tx
			go s.write(tx)
		}
	}
}

func (s *SP2p) writeTx(msg *KMsg) {
	s.txWC <- msg
}

func (s *SP2p) write(msg *KMsg) {
	if msg.Addr == "" {
		msg.Addr = s.GetAddr()
	}
	if msg.ID == "" {
		msg.ID = s.tab.selfNode.ID.ToHex()
	}
	if msg.RID == "" {
		msg.RID = <-cfg.uuidC
	}
	if msg.Version == "" {
		msg.Version = cfg.Version
	}
	if msg.TAddr == "" {
		GetLog().Error("target node addr is nonexistent")
		return
	}
	if msg.TID == "" {
		GetLog().Error("target node id is nonexistent")
		return
	}

	addr, err := net.ResolveUDPAddr("udp", msg.TAddr)
	if err != nil {
		GetLog().Error("ResolveUDPAddr error", "err", err)
		return
	}

	if _, err := s.rconn.WriteToUDP(msg.Dumps(), addr); err != nil {
		GetLog().Error("WriteToUDP error", "err", err)
		return
	}
}

func (s *SP2p) pingN() {
	for _, n := range s.tab.FindRandomNodes(cfg.PingNodeNum) {
		s.writeTx(&KMsg{TAddr: n.AddrString(), TID: n.ID.ToHex(), Data: &PingReq{}})
	}
}

func (s *SP2p) findN() {
	for _, b := range s.tab.buckets {
		if b == nil || b.size() == 0 {
			continue
		}
		n := b.Random()
		s.writeTx(&KMsg{TAddr: n.AddrString(), TID: n.ID.ToHex(), Data: &FindNodeReq{N: cfg.FindNodeNUm}})
	}
}

func (s *SP2p) accept() {
	kb := NewKBuffer()
	logger := GetLog()
	for {
		buf := make([]byte, cfg.MaxBufLen)
		n, err := s.conn.Read(buf)
		if err != nil {
			if strings.Contains(err.Error(), "timeout") {
				GetLog().Error("timeout", "err", err)
			} else if err == io.EOF {
				GetLog().Error("udp read eof ", "err", err)
				break
			} else if err != nil {
				GetLog().Error("udp read error ", "err", err)
			}
			time.Sleep(time.Second * 2)
			continue
		}

		messages := kb.Next(buf[:n])
		if messages == nil {
			continue
		}

		for _, m := range messages {
			if m == nil || len(m) == 0 {
				continue
			}

			msg := &KMsg{}
			if err := msg.Decode(m); err != nil {
				logger.Error("kmsg decode error", "err", err.Error(), "method", "sp2p.accept")
				continue
			}

			s.txRC <- msg
		}
	}
}
