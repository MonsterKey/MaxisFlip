package tool

import (
	"flipBot/lib/api"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"math/big"
	"time"
)

// CallData structure
type CallData struct {
	Time          int64
	ERCToken      string
	BetAmount     string
	Bet           float64
	BetFace       string
	RaffleAddress string
}

// FlipOperation performs the coin flip operation
func FlipOperation(betAmount float64, address, sk string) (interface{}, error) {
	calldata := CallData{
		Time:          time.Now().Unix(),
		ERCToken:      "0x0000000000000000000000000000000000000000",
		BetAmount:     new(big.Float).Mul(big.NewFloat(betAmount), big.NewFloat(1e18)).Text('f', 0), // Multiply by 10^18
		Bet:           betAmount,
		BetFace:       fmt.Sprintf("%d", GetRandomOneOrTwo()), // Assuming this function returns "1" or "2"
		RaffleAddress: "0xe9536b6a37ec1f7b3b3d76c2761a03d7e72ccaa1",
	}
	//fmt.Println("calldata:", calldata)

	message := solsha3.SoliditySHA3(
		solsha3.Address(address),
		solsha3.Address(calldata.ERCToken),
		solsha3.Uint256(calldata.BetAmount),
		solsha3.Uint64(calldata.Time),
		solsha3.Address(calldata.RaffleAddress),
		solsha3.Uint256("137"),
	)
	signature, err := SignMessageWeb3(common.BytesToHash(message).String()[2:], true, sk)
	if err != nil {
		fmt.Println("Error signing message:", err)
		return nil, err
	}

	resp, err := api.FlipPayment(api.FlipPaymentReqParam{
		Time:          calldata.Time,
		ErcToken:      calldata.ERCToken,
		BetAmount:     calldata.BetAmount,
		Bet:           calldata.Bet,
		BetFace:       calldata.BetFace,
		RaffleAddress: calldata.RaffleAddress,

		Signer:        address,
		Signature:     signature,
	})
	if err != nil {
		return nil, err
	}
	if resp.Code != 200 {
		return nil, fmt.Errorf("error from server: %s", resp.Msg)
	}
	//fmt.Println("response:", resp.Data)

	var historyResult interface{}
	for {
		time.Sleep(3 * time.Second) // Polling interval
		history, err := api.GetCoinFlipHistoryList(api.FlipHistoryListReqParam{
			Owner: address,
			Time:  calldata.Time,
			Token: "0x0000000000000000000000000000000000000000",
		})
		if err != nil {
			fmt.Println("getCoinFlipHistoryList error:", err)
			continue
		}
		//fmt.Println("history:", history)
		if len(history.Data) > 0 {
			historyResult = history.Data[0]
			break
		}
	}

	return historyResult, nil
}
