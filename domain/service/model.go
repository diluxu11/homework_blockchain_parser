package service

const (
	Version = "2.0"

	ethBlockNumberMethod      = "eth_blockNumber"
	ethUninstallFilterMethod  = "eth_uninstallFilter"
	ethGetFilterChangesMethod = "eth_getFilterChanges"
)

type CommonRequest struct {
	ID      int64       `json:"id"`
	JSONRpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CommonResponse struct {
	ID     int    `json:"id"`
	Result string `json:"result"`
	Error  *Error `json:"error"`
}
