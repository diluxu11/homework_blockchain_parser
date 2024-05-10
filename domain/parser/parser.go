package parser

import (
	"awesomeProject/dal/redis"
	"awesomeProject/model"
	"context"
)

type IParser interface {
	// last parsed block
	GetCurrentBlock() int
	// add address to observer
	Subscribe(address string)
	// check subscription status of address
	CheckSubscription(address string) bool
	// list of inbound or outbound transactions for an address
	GetTransactions(address string) []*model.Transaction
}

type Parser struct {
	ctx context.Context
}

func (n *Parser) GetCurrentBlock() int {
	return redis.GetCurrentScannedBlock()
}

func (n *Parser) Subscribe(address string) {
	redis.AddAddress2Subscription(address)
}

func (n *Parser) CheckSubscription(address string) bool {
	return redis.CheckSubscription(address)
}

func (n *Parser) GetTransactions(address string) []*model.Transaction {
	return redis.GetTransactionListByAddress(address)
}
