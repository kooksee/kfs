package packets

const (
	KVSetReqT = byte(0x4)
	KVSetReqS = "kv set req"

	KVGetReqT = byte(0x5)
	KVGetReqS = "kv get req"

	KVGetRespT = byte(0x6)
	KVGetRespS = "kv get resp"

	ChatReqT = byte(0x7)
	ChatReqS = "chat req"

	ChatRespT = byte(0x8)
	ChatRespS = "chat resp"
)
