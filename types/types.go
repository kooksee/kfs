package types

type RPCRequest struct {
	ID        string                 `json:"id"`
	Method    string                 `json:"method"`
	Signature string                 `json:"sign,omitempty"`
	Params    map[string]interface{} `json:"params,omitempty"`
}

type RPCResponse struct {
	Code int    `json:"code"`
	ID   string `json:"id"`
	Msg  string `json:"msg,omitempty"`
	Data string `json:"data,omitempty"`
}
