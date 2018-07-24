package sp2p

import (
	"time"

	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/kooksee/kdb"
	"encoding/hex"
	"errors"
)

var BucketPrefix = []byte("bkt")

type bucket struct {
	peers *arraylist.List
	h     *kdb.KHash
}

func newBuckets() *bucket {
	return &bucket{
		peers: arraylist.New(),
		h:     GetDb().KHash(BucketPrefix),
	}
}

func (b *bucket) updateNodes(nodes ... *Node) {
	for _, n := range nodes {
		n.updateAt = time.Now()
		b.addNodes(n)
	}
}

// addNode add node to bucket, if bucket is full, will remove an old one
func (b *bucket) addNodes(nodes ... *Node) {

	logger := GetLog()

	// 把最活跃的放到最前面,然后移除最不活跃的
	if err := b.h.WithTx(func(k *kdb.KHBatch) error {
		for _, node := range nodes {
			logger.Info("add node", "node", node.String())
			b.peers.Add(node)
			if err := k.Set(NodesBackupKey(node.ID.Bytes()), []byte(node.String())); err != nil {
				logger.Error("add peer error", "err", err)
				continue
			}
		}

		b.peers.Sort(func(a, b interface{}) int { return int(b.(*Node).updateAt.Sub(a.(*Node).updateAt)) })
		size := b.peers.Size()
		if size < cfg.BucketSize {
			return errors.New("")
		}

		for i := cfg.BucketSize; i < size; i++ {
			val, e := b.peers.Get(i)
			if !e {
				continue
			}
			b.peers.Remove(i)
			if err := k.MDel(NodesBackupKey(val.(*Node).ID.Bytes())); err != nil {
				logger.Error("delete peer error", "err", err)
				continue
			}
		}
		return nil
	}); err != nil {

	}

}

// findNode check if the bucket already have this node, if so, return its index, otherwise, return -1
func (b *bucket) findNode(node *Node) int {
	return b.peers.IndexOf(node)
}

func (b *bucket) Random() *Node {
	if b.size() == 0 {
		return nil
	}

	val, _ := b.peers.Get(int(randUint(uint32(b.size()))))
	return val.(*Node)
}

func (b *bucket) deleteNodes(targets ... Hash) {
	if err := b.h.WithTx(func(k *kdb.KHBatch) error {
		for _, node := range targets {
			if a := b.peers.IndexOf(node); a != -1 {
				val, bl := b.peers.Get(a)
				if !bl {
					continue
				}
				if err := k.MDel(NodesBackupKey(val.(*Node).ID.Bytes())); err != nil {
					GetLog().Error("deleteNodes error", "err", err)
					continue
				}
				GetLog().Info("delete node: %s", hex.EncodeToString(node.Bytes()))
				b.peers.Remove(a)
			}
		}
		return nil
	}); err != nil {
		GetLog().Error("update peer", "err", err)
	}
}

func (b *bucket) size() int {
	return b.peers.Size()
}
