package service

import "awesomeProject/model"

const (
	Version = "2.0"

	ethBlockNumberMethod      = "eth_blockNumber"
	ethGetBlockByNumberMethod = "eth_getBlockByNumber"

	URL = "https://cloudflare-eth.com"
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

type Block struct {
	Number       string               `json:"number"`
	Hash         string               `json:"hash"`
	Transactions []*model.Transaction `json:"transactions"`
}

type GetBlockByNumberResponse struct {
	ID     int    `json:"id"`
	Result *Block `json:"result"`
	Error  *Error `json:"error"`
}
