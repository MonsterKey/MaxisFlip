package api

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type FlipHistoryListReqParam struct {
	Owner string `json:"owner"`
	Time  int64  `json:"time"`
	Token string `json:"token"`
}

type FlipPaymentReqParam struct {
	Time          int64   `json:"time"`
	ErcToken      string  `json:"ercToken"`
	BetAmount     string  `json:"betAmount"`
	Bet           float64 `json:"bet"`
	BetFace       string  `json:"betFace"`
	RaffleAddress string  `json:"raffleAddress"`
	Signer        string  `json:"signer"`
	Signature     string  `json:"signature"`
}

type ResponseData struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type ResponseCoinFlipHistoryList struct {
	Code int `json:"code"`
	Data []struct {
		ID          string `json:"_id"`
		Chain       string `json:"chain"`
		Owner       string `json:"owner"`
		Time        int    `json:"time"`
		V           int    `json:"__v"`
		Amount      string `json:"amount"`
		BetFace     string `json:"betFace"`
		CreateAt    int64  `json:"createAt"`
		FeeAmount   int64  `json:"feeAmount"`
		FeeReceiver string `json:"feeReceiver"`
		Token       string `json:"token"`
		UpdateAt    int64  `json:"updateAt"`
		WinAmount   int    `json:"winAmount"`
		WinFace     string `json:"winFace"`
		WinTimes    int    `json:"winTimes"`
		Won         bool   `json:"won"`
	} `json:"data"`
}

var (
	Host = "https://maxis.gg"
)

// FlipPayment coinFlip/flip api
func FlipPayment(param FlipPaymentReqParam) (ResponseData, error) {
	client := resty.New()
	var responseData ResponseData
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(param).
		SetResult(&responseData).
		Post(fmt.Sprintf("%s/api/coinFlip/flip", Host))

	if err != nil {
		return ResponseData{}, err
	}

	return responseData, nil
}

// GetCoinFlipHistoryList coinFlip/history api
func GetCoinFlipHistoryList(param FlipHistoryListReqParam) (ResponseCoinFlipHistoryList, error) {
	client := resty.New()
	var responseData ResponseCoinFlipHistoryList
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(param).
		SetResult(&responseData).
		Post(fmt.Sprintf("%s/api/coinFlip/history", Host))

	if err != nil {
		return ResponseCoinFlipHistoryList{}, err
	}

	return responseData, nil
}