package tool

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type Config struct {
	BetSelectProbability []BetSelectProbability
	Wallets              []Task
}

type BetSelectProbability struct {
	BetSelect float64 `json:"betSelect"`
	Prob      float64 `json:"prob"`
}

type Task struct {
	Pk        string  `json:"pk"`
	Sk        string  `json:"sk"`
	PayAmount float64 `json:"payAmount"`
}

func hashMessage(hexMessage string) []byte {
	messageBytes, err := hex.DecodeString(hexMessage)
	if err != nil {
		panic(any(err))
	}

	prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(messageBytes))
	fullMessage := prefix + string(messageBytes)
	hash := crypto.Keccak256Hash([]byte(fullMessage))
	return hash.Bytes()
}

func isValidAddress(address string) bool {
	return common.IsHexAddress(address)
}

// SignMessageWeb3 signs a message using the provided private key
func SignMessageWeb3(message string, msg16 bool, sk string) (string, error) {
	privateKey, err := crypto.HexToECDSA(sk)
	if err != nil {
		return "", err
	}

	var messageBytes []byte
	if msg16 {
		messageBytes = hashMessage(message)
	} else {
		message = fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
		messageHash := crypto.Keccak256Hash([]byte(message))
		messageBytes = messageHash.Bytes()
	}

	// Sign the hash of the message
	signature, err := crypto.Sign(messageBytes, privateKey)
	if err != nil {
		return "", err
	}
	// Add Ethereum signature prefix
	if signature[64] < 27 {
		signature[64] += 27
	}

	return hexutil.Encode(signature), nil
}

// GetRandomOneOrTwo returns a random number between 1 and 2
func GetRandomOneOrTwo() int {
	return rand.Intn(2) + 1
}

// RandomSelect selects an element from betSelects based on the given probabilities
func RandomSelect(betSelects []float64, probabilities []float64) float64 {
	var total float64
	cumulative := make([]float64, len(probabilities))

	// Compute the cumulative probability array
	for i, prob := range probabilities {
		total += prob
		cumulative[i] = total
	}

	// Generate a random number in the range [0, total)
	rand.Seed(time.Now().UnixNano())
	random := rand.Float64() * total

	// Select the bet based on the random number
	for i, cum := range cumulative {
		if random < cum {
			return betSelects[i]
		}
	}

	return 0 // return an empty string if no selection is made
}

// ReadConfigInfo reads a JSON file and returns the parsed data
func ReadConfigInfo(path string) (Config, error) {
	// read file
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return Config{}, err
	}

	// parse JSON
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return Config{}, err
	}

	return config, nil
}

// GetFromObjBetSelectProbability returns the BetSelect and Prob values from a list of BetSelectProbability objects
func GetFromObjBetSelectProbability(bets []BetSelectProbability) ([]float64, []float64) {
	var (
		betSelects    []float64
		probabilities []float64
	)
	for _, probability := range bets {
		betSelects = append(betSelects, probability.BetSelect)
		probabilities = append(probabilities, probability.Prob)
	}

	return betSelects, probabilities
}

// CheckPrivateKeyWithPublicKey checks if the provided private key corresponds to the expected public key
func CheckPrivateKeyWithPublicKey(privateKeyHex, expectedPublicKeyHex string) bool {
	if !isValidAddress(expectedPublicKeyHex) {
		fmt.Println("Invalid address:", expectedPublicKeyHex)
		return false
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		fmt.Println("Invalid private key:", err)
		return false
	}

	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("Error casting public key to ECDSA")
		return false
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	if strings.ToLower(address) == strings.ToLower(expectedPublicKeyHex) {
		return true
	} else {
		fmt.Printf("Mismatched public key: Got %s, Expected %s\n", address, expectedPublicKeyHex)
		return false
	}
}
