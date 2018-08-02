package packets

import "github.com/kooksee/sp2p"

/*
实现p2p的聊天方式
 */

type chatReq struct {
	Method string                 `json:"method"`
	Params map[string]interface{} `json:"params"`
}

func (t *chatReq) T() byte        { return chatReqT }
func (t *chatReq) String() string { return chatReqS }
func (t *chatReq) OnHandle(p sp2p.ISP2P, msg *sp2p.KMsg) {
	// 获得聊天的结果

	switch t.Method {
	default:
		logger.Error("方法不正确")
	case "":
	}
}

type chatResp struct {
	Method string `json:"method"`
	Data   string `json:"data"`
}

func (t *chatResp) T() byte        { return chatRespT }
func (t *chatResp) String() string { return chatRespS }
func (t *chatResp) OnHandle(p sp2p.ISP2P, msg *sp2p.KMsg) {
	// 获得聊天的结果

	switch t.Method {
	default:
		logger.Error("方法不正确")
	case "":
	}
}
