package sp2p

import (
	"errors"
)

type KMsg struct {
	Version string   `json:"version,omitempty"`
	RID     string   `json:"rid,omitempty"`
	ID      string   `json:"id,omitempty"`
	Addr    string   `json:"addr,omitempty"`
	TAddr   string   `json:"taddr,omitempty"`
	TID     string   `json:"tid,omitempty"`
	Data    IMessage `json:"data,omitempty"`
}

func (t *KMsg) Decode(msg []byte) error {
	dt := msg[0]
	if !hm.Contain(dt) {
		return errors.New(Fmt("kmsg type %s is nonexistent", dt))
	}

	t.Data = hm.GetHandler(dt)
	return json.Unmarshal(msg[1:], t)
}

func (t *KMsg) Dumps() []byte {
	d, _ := json.Marshal(t)
	return append([]byte{t.Data.T()}, append(d, "\n"...)...)
}
