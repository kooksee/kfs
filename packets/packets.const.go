package packets

const (
	kVSetReqT = byte(0x4)
	kVSetReqS = "kv set req"

	kVGetReqT = byte(0x5)
	kVGetReqS = "kv get req"

	kVGetRespT = byte(0x6)
	kVGetRespS = "kv get resp"

	chatReqT = byte(0x7)
	chatReqS = "chat req"

	chatRespT = byte(0x8)
	chatRespS = "chat resp"
)
