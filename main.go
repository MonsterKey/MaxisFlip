package main

import (
	"encoding/json"
	"flag"
	tool "flipBot/utils"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	// Command line flags
	pkPrefix     = flag.String("pk", "", "wallet public key")
	skPrefix     = flag.String("sk", "", "wallet private key")
	amountPrefix = flag.Float64("amount", 0.0, "total pay amount")
	betPrefix    = flag.Float64("bet", 0.0, "bet select")
	configPath   = flag.String("config", "./config.json", "config file path")

	// Mode Default config
	config tool.Config
)

func init() {
	flag.Parse()

	if *pkPrefix != "" && *skPrefix != "" && *amountPrefix != 0.0 && *betPrefix != 0.0 {
		config = tool.Config{
			BetSelectProbability: []tool.BetSelectProbability{
				{
					BetSelect: *betPrefix,
					Prob:      100,
				},
			},
			Wallets: []tool.Task{
				{
					Pk:        *pkPrefix,
					Sk:        *skPrefix,
					PayAmount: *amountPrefix,
				},
			},
		}
	}
}

func startRun(pk, sk string, totalPayAmount float64, bets []tool.BetSelectProbability) {
	currentPayAmount := 0.0
	betSelects, probabilities := tool.GetFromObjBetSelectProbability(bets)

	// check pk and sk
	if !tool.CheckPrivateKeyWithPublicKey(sk, pk) {
		fmt.Println("Invalid Private Key and Public Key")
		os.Exit(1)
	}

	// Run until the total pay amount is reached
	for {
		betAmount := tool.RandomSelect(betSelects, probabilities)
		fmt.Printf("%s - pay(%.3f) - currentTotal(%.3f)\n", pk, betAmount, currentPayAmount)

		_, err := tool.FlipOperation(betAmount, pk, sk)
		if err != nil {
			fmt.Println("Error:", err)
			if strings.Contains(err.Error(), "Insufficient") {
				break
			}
		} else {
			currentPayAmount += betAmount
			if currentPayAmount >= totalPayAmount {
				break
			}
		}

		sleepDuration := time.Duration(rand.Intn(9000) + 1000) // Random sleep between 1000 and 10000 milliseconds
		time.Sleep(sleepDuration)
	}

	fmt.Printf("%s ended - total: %.3f\n", pk, currentPayAmount)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("version: flip bot v0.1.0")
	if len(config.Wallets) <= 0 {
		config, _ = tool.ReadConfigInfo(*configPath)
	}

	runTasker := config.Wallets
	formattedJSON, _ := json.MarshalIndent(runTasker, "", "  ")
	fmt.Println("runTasker:", string(formattedJSON))

	for _, item := range runTasker {
		go func(task tool.Task) {
			startRun(strings.ToLower(task.Pk), task.Sk, task.PayAmount, config.BetSelectProbability)
		}(item)

		// sleep between 20 and 30 seconds
		sleepDuration := time.Duration(rand.Intn(10)+20) * time.Second // Random sleep between 20 and 30 seconds
		time.Sleep(sleepDuration)
	}

	// wait for goroutines to finish
	_, _ = fmt.Scanln()
}
