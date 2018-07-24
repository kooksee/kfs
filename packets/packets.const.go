package sp2p

const (
	KVSetReqT = byte(0x4)
	KVSetReqS = "kv set req"

	KVGetReqT = byte(0x5)
	KVGetReqS = "kv get req"

	KVGetRespT = byte(0x6)
	KVGetRespS = "kv get resp"

	GKVSetReqT = byte(0x7)
	GKVSetReqS = "gossip kv set req"

	GKVGetReqT = byte(0x8)
	GKVGetReqS = "gossip kv get req"

	GKVGetRespT = byte(0x9)
	GKVGetRespS = "gossip kv get resp"
)
