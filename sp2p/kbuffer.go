package sp2p

import (
	"bytes"
	"sync"
)

func NewKBuffer() *KBuffer {
	return &KBuffer{dmt: []byte{'\n'}}
}

type KBuffer struct {
	buf []byte
	dmt []byte
	sync.RWMutex
}

func (t *KBuffer) SetDmt(dmt []byte) *KBuffer {
	t.dmt = dmt
	return t
}

func (t *KBuffer) Next(b []byte) [][]byte {
	t.Lock()
	defer t.Unlock()

	if b == nil {
		return nil
	}

	t.buf = append(t.buf, b...)

	if len(t.buf) < 1 {
		return nil
	}

	if !bytes.Contains(t.buf, t.dmt) {
		return nil
	}

	d := bytes.Split(t.buf, t.dmt)
	if len(d) < 1 {
		return nil
	}

	t.buf = d[len(d)-1]
	return d[:len(d)-1]
}
