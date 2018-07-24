package sp2p

import (
	"sync"
	"reflect"
)

var (
	hmOnce sync.Once
	hm     *HandleManager
)

func GetHManager() *HandleManager {
	hmOnce.Do(func() {
		hm = &HandleManager{hmap: make(map[byte]reflect.Type)}
	})
	return hm
}

type HandleManager struct {
	hmap map[byte]reflect.Type
}

func (h *HandleManager) HandleTypes() []byte {
	a := make([]byte, 0)
	for k := range h.hmap {
		a = append(a, k)
	}
	return a
}

func (h *HandleManager) Registry(handlers ... interface{}) {
	for _, handler := range handlers {

		h1 := reflect.TypeOf(handler)
		h3 := reflect.New(h1).Interface().(IMessage)

		name := h3.T()
		if h.Contain(name) {
			GetLog().Error("registry error handler exist", "type", name, "desc", h3.String())
			panic("")
		}
		h.hmap[name] = h1
	}
}

func (h *HandleManager) Contain(name byte) bool {
	_, ok := h.hmap[name]
	return ok
}

func (h *HandleManager) GetHandler(name byte) IMessage {
	h1 := h.hmap[name]
	h2 := reflect.New(h1)
	return h2.Interface().(IMessage)
}
