package handler

import (
	"awesomeProject/dal/redis"
	"awesomeProject/domain/service"
	"context"
	"fmt"
	"strconv"
)

func Scan() {
	ctx := context.Background()

	if !CheckRunnable(ctx) {
		return
	}

	current := redis.GetCurrentScannedBlock()
	transactionList, err := service.GetTransactionListByBlockNum(ctx, fmt.Sprintf("0x%d", strconv.FormatInt(int64(current), 16)))
	if err != nil {
		return
	}

	for _, v := range transactionList {
		redis.AddTransaction(v)
	}

	redis.UpdateCurrentScannedBlock(current + 1)
}

// CheckRunnable To check whether current node is already the latest node
func CheckRunnable(ctx context.Context) bool {
	current := redis.GetCurrentScannedBlock()
	if current == 0 {
		return true
	}

	blockNum, err := service.GetLatestBlockNumber(ctx)
	if err != nil {
		return false
	}

	num, err := strconv.ParseInt(blockNum[2:], 16, 64)
	if err != nil {
		return false
	}

	if int(num) > current {
		return true
	}

	return false
}
