package parser

import (
	"awesomeProject/model"
	"context"
)

type IParser interface {
	// last parsed block
	GetCurrentBlock() int
	// add address to observer
	Subscribe(address string) bool
	// list of inbound or outbound transactions for an address
	GetTransactions(address string) []model.Transaction
}

type Parser struct {
	ctx context.Context
}

func (n *Parser) GetCurrentBlock() int {
	return 0
}

func (n *Parser) Subscribe(address string) bool {
	return true
}

func (n *Parser) GetTransactions(address string) []model.Transaction {
	return nil
}

type NotificationParser struct {
	Parser
}

func (n *NotificationParser) GetCurrentBlock() int {
	return 0
}

func (n *NotificationParser) Subscribe(address string) bool {
	return true
}

func (n *NotificationParser) GetTransactions(address string) []model.Transaction {
	return nil
}