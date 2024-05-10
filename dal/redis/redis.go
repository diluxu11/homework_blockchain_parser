package redis

import (
	"awesomeProject/model"
	"sync"
)

var (
	currentScannedBlock CurrentScannedBlock
	subscription        Subscription
	transactionMapData  TransactionMap
)

type CurrentScannedBlock struct {
	lock sync.RWMutex
	data int
}

type Subscription struct {
	lock sync.RWMutex
	data map[string]bool
}

type TransactionMap struct {
	lock sync.RWMutex
	data map[string][]*model.Transaction
}

func Init() {
	// use in app memory to mock redis
	currentScannedBlock = CurrentScannedBlock{
		data: 0,
	}
	subscription = Subscription{
		data: make(map[string]bool),
	}
	transactionMapData = TransactionMap{
		data: make(map[string][]*model.Transaction),
	}
}

// GetCurrentScannedBlock GetCurrentBlockNumber
func GetCurrentScannedBlock() int {
	currentScannedBlock.lock.Lock()
	defer currentScannedBlock.lock.Unlock()
	return currentScannedBlock.data
}

// UpdateCurrentScannedBlock Update Scanned Block
func UpdateCurrentScannedBlock(block int) {
	currentScannedBlock.lock.Lock()
	defer currentScannedBlock.lock.Unlock()
	currentScannedBlock.data = block
}

// GetTransactionListByAddress GetTransactionList
func GetTransactionListByAddress(address string) []*model.Transaction {
	transactionMapData.lock.Lock()
	defer transactionMapData.lock.Unlock()
	if _, ok := transactionMapData.data[address]; !ok {
		return nil
	}

	return transactionMapData.data[address]
}

// AddTransaction Add Transaction to map
func AddTransaction(t *model.Transaction) {
	transactionMapData.lock.Lock()
	defer transactionMapData.lock.Unlock()

	if list, ok := transactionMapData.data[t.FromAddress]; ok {
		list = append(list, t)
	} else {
		transactionMapData.data[t.FromAddress] = []*model.Transaction{t}
	}

	if list, ok := transactionMapData.data[t.ToAddress]; ok {
		list = append(list, t)
	} else {
		transactionMapData.data[t.ToAddress] = []*model.Transaction{t}
	}
}

// AddAddress2Subscription Subscribe a new address
func AddAddress2Subscription(address string) {
	subscription.lock.Lock()
	defer subscription.lock.Unlock()
	subscription.data[address] = true
}

// CheckSubscription Check subscription status of address
func CheckSubscription(address string) bool {
	subscription.lock.Lock()
	defer subscription.lock.Unlock()
	if _, ok := subscription.data[address]; !ok {
		return false
	}
	return subscription.data[address]
}
