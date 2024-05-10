package service

import (
	"awesomeProject/model"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Timeout: time.Second,
}

// GetLatestBlockNumber send a request to get Latest block number
func GetLatestBlockNumber(ctx context.Context) (string, error) {
	data, err := json.Marshal(&CommonRequest{
		ID:      1,
		JSONRpc: Version,
		Method:  ethBlockNumberMethod,
		Params:  nil,
	})

	if err != nil {
		return "", errors.New("invalid Params")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, URL, bytes.NewReader(data))
	if err != nil {
		return "", errors.New(fmt.Sprintf("GetLatestBlockNumber err:%v\n", err.Error()))
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := httpClient.Do(req)

	if err != nil {
		return "", errors.New(fmt.Sprintf("GetLatestBlockNumber err:%v\n", err.Error()))
	}

	if res.StatusCode == http.StatusOK {
		return "", errors.New(fmt.Sprintf("GetLatestBlockNumber statusCode = %d", res.StatusCode))
	}

	resp := CommonResponse{}
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return "", fmt.Errorf("GetLatestBlockNumber decoding response body err: %v", err)
	}

	return resp.Result, nil
}

// GetTransactionListByBlockNum Get TransactionList by BlockNum Scanned
func GetTransactionListByBlockNum(ctx context.Context, blockNum string) ([]*model.Transaction, error) {
	data, err := json.Marshal(&CommonRequest{
		ID:      1,
		JSONRpc: Version,
		Method:  ethGetBlockByNumberMethod,
		Params:  []string{blockNum},
	})

	if err != nil {
		return nil, errors.New("invalid Params")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, URL, bytes.NewReader(data))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("GetTransactionListByBlockNum err:%v\n", err.Error()))
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := httpClient.Do(req)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("GetTransactionListByBlockNum err:%v\n", err.Error()))
	}

	if res.StatusCode == http.StatusOK {
		return nil, errors.New(fmt.Sprintf("GetTransactionListByBlockNum statusCode = %d", res.StatusCode))
	}

	resp := GetBlockByNumberResponse{}
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, fmt.Errorf("GetTransactionListByBlockNum decoding response body err: %v", err)
	}

	return resp.Result.Transactions, nil
}
