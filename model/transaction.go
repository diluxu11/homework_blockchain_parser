package model

type Transaction struct {
	Hash        string `json:"transactionHash"`
	FromAddress string `json:"from_address"`
	ToAddress   string `json:"to_address"`
	BlockNumber string `json:"blockNumber"`
	BlockHash   string `json:"blockHash"`
}
