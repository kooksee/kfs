package sp2p

import (
	"time"
	"github.com/inconshreveable/log15"
	"github.com/kooksee/kdb"
	"os"
	"path/filepath"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

var cfg *KConfig

type KConfig struct {
	// 接收数据的最大缓存区
	MaxBufLen int

	// ntp服务器检测超时次数
	NtpFailureThreshold int
	//在重复NTP警告之前需要经过的最短时间
	NtpWarningCooldown time.Duration
	// ntpPool is the NTP server to query for the current time
	NtpPool string
	// Number of measurements to do against the NTP server
	NtpChecks int
	// Allowed clock drift before warning user
	DriftThreshold time.Duration

	PingTick     *time.Ticker
	FindNodeTick *time.Ticker
	NtpTick      *time.Ticker
	RelayTick    *time.Ticker

	// Kademlia concurrency factor
	Alpha int
	// 节点响应的数量
	NodeResponseNumber int
	// 节点广播的数量
	NodeBroadcastNumber int
	// 节点分区存储的数量
	NodePartitionNumber int

	PingNodeNum int
	FindNodeNUm int

	// 节点ID长度
	HashBits int

	NodesBackupKey string

	BucketSize int

	MaxNodeSize int
	MinNodeSize int
	Version     string

	StoreAckNum int

	Adds []string

	Seeds []string
	KvKey []byte

	KeyStore *keystore.KeyStore
	uuidC    chan string
	db       *kdb.KDB
	l        log15.Logger
}

func (t *KConfig) InitLog(l ... log15.Logger) *KConfig {
	if len(l) != 0 {
		t.l = l[0].New("package", "sp2p")
	} else {
		t.l = log15.New("package", "sp2p")
		t.l.SetHandler(log15.LvlFilterHandler(log15.LvlDebug, log15.StreamHandler(os.Stdout, log15.TerminalFormat())))
	}
	return t
}

func (t *KConfig) InitDb(db ... *kdb.KDB) *KConfig {
	if len(db) != 0 {
		t.db = db[0]
	} else {
		kdb.InitKdb(filepath.Join("kdata", "db"))
		t.db = kdb.GetKdb()
	}
	return t
}

func GetLog() log15.Logger {
	if GetCfg().l == nil {
		panic("please init sp2p log")
	}
	return GetCfg().l
}

func GetDb() *kdb.KDB {
	if GetCfg().db == nil {
		GetLog().Error("please init sp2p db")
		panic("")
	}
	return GetCfg().db
}

func GetCfg() *KConfig {
	if cfg == nil {
		panic("please init sp2p config")
	}
	return cfg
}

func InitCfg() *KConfig {
	cfg = &KConfig{
		MaxBufLen:           1024 * 16,
		NtpFailureThreshold: 32,
		NtpWarningCooldown:  10 * time.Minute,
		NtpPool:             "pool.ntp.org",
		NtpChecks:           3,
		DriftThreshold:      10 * time.Second,
		Alpha:               3,
		NodeResponseNumber:  8,
		NodeBroadcastNumber: 16,
		NodePartitionNumber: 8,
		HashBits:            len(Hash{}) * 8,
		PingNodeNum:         8,
		FindNodeNUm:         20,

		Adds: []string{"127.0.0.1:8080"},

		NodesBackupKey: "nbk:",

		PingTick:     time.NewTicker(10 * time.Minute),
		FindNodeTick: time.NewTicker(1 * time.Hour),
		NtpTick:      time.NewTicker(10 * time.Minute),
		RelayTick:    time.NewTicker(time.Minute),

		MaxNodeSize: 2000,
		MinNodeSize: 100,
		Version:     "1.0.0",

		BucketSize:  16,
		StoreAckNum: 2,

		KvKey: []byte("kv:"),

		uuidC: make(chan string, 500),
	}

	return cfg
}
