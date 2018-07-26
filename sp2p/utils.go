package sp2p

import (
	"fmt"
	"math/rand"
	"net"
	"time"
	"strings"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
)

func Errs(err ... string) string {
	return strings.Join(err, "->")
}

var Fmt = fmt.Sprintf

// DistCmp compares the distances a->target and b->target.
// Returns -1 if a is closer to target, 1 if b is closer to target
// and 0 if they are equal.
func DistCmp(target, a, b Hash) int {
	for i := range target {
		da := a[i] ^ target[i]
		db := b[i] ^ target[i]
		if da > db {
			return 1
		} else if da < db {
			return -1
		}
	}
	return 0
}

func Expired(ts int64) bool {
	return time.Unix(ts, 0).Before(time.Now())
}

func TimeAdd(ts time.Duration) time.Time {
	return time.Now().Add(ts)
}

func If(cond bool, trueVal, falseVal interface{}) interface{} {
	if cond {
		return trueVal
	}
	return falseVal
}

// logdist returns the logarithmic distance between a and b, log2(a ^ b).
func Logdist(a, b Hash) int {
	lz := 0
	for i := range a {
		x := a[i] ^ b[i]
		if x == 0 {
			lz += 8
		} else {
			lz += lzcount[x]
			break
		}
	}
	return len(a)*8 - lz
}

// hashAtDistance returns a random hash such that logdist(a, b) == n
func HashAtDistance(a Hash, n int) (b Hash) {
	if n == 0 {
		return a
	}
	// flip bit at position n, fill the rest with random bits
	b = a
	pos := len(a) - n/8 - 1
	bit := byte(0x01) << (byte(n%8) - 1)
	if bit == 0 {
		pos++
		bit = 0x80
	}
	b[pos] = a[pos]&^bit | ^a[pos]&bit // TODO: randomize end bits
	for i := pos + 1; i < len(a); i++ {
		b[i] = byte(rand.Intn(255))
	}
	return b
}

func NodeFromKMsg(msg *KMsg) (*Node, error) {
	nid, err := HexID(msg.ID)
	if err != nil {
		return nil, err
	}
	addr, err := net.ResolveUDPAddr("udp", msg.Addr)
	if err != nil {
		return nil, err
	}
	return NewNode(nid, addr.IP, uint16(addr.Port)), nil
}

func MustNotErr(err error) {
	if err == nil {
		return
	}
	GetLog().Error("MustNotErr", "err", err)
	panic(err.Error())
}

func NodesBackupKey(k []byte) []byte {
	return append([]byte(cfg.NodesBackupKey), k...)
}

// table of leading zero counts for bytes [0..255]
var lzcount = [256]int{
	8, 7, 6, 6, 5, 5, 5, 5,
	4, 4, 4, 4, 4, 4, 4, 4,
	3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3,
	2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2,
	1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
}

// 生成私钥
func GenPrivateKey() *ecdsa.PrivateKey {
	pkv, err := crypto.GenerateKey()
	if err != nil {
		panic(err.Error())
	}
	return pkv
}

// 加载私钥
func LoadPriv(p []byte) *ecdsa.PrivateKey {
	pkv, err := crypto.ToECDSA(p)
	if err != nil {
		panic(err.Error())
	}
	return pkv
}
