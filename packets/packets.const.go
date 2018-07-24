package packets

const (
	KVSetReqT = byte(0x4)
	KVSetReqS = "kv set req"

	KVGetReqT = byte(0x5)
	KVGetReqS = "kv get req"

	KVGetRespT = byte(0x6)
	KVGetRespS = "kv get resp"
)
