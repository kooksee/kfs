package packets

import "github.com/kooksee/sp2p"

/*
实现p2p的聊天方式
 */

type ChatReq struct {
	Method string                 `json:"method"`
	Params map[string]interface{} `json:"params"`
}

func (t *ChatReq) T() byte        { return ChatReqT }
func (t *ChatReq) String() string { return ChatReqS }
func (t *ChatReq) OnHandle(p sp2p.ISP2P, msg *sp2p.KMsg) {
	// 获得聊天的结果

	switch t.Method {
	default:
		logger.Error("方法不正确")
	case "":
	}
}

type ChatResp struct {
	Method string `json:"method"`
	Data   string `json:"data"`
}

func (t *ChatResp) T() byte        { return ChatRespT }
func (t *ChatResp) String() string { return ChatRespS }
func (t *ChatResp) OnHandle(p sp2p.ISP2P, msg *sp2p.KMsg) {
	// 获得聊天的结果

	switch t.Method {
	default:
		logger.Error("方法不正确")
	case "":
	}
}
