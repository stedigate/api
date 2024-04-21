package ethereum

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nanmu42/etherscan-api"
	"github.com/redis/rueidis"
	"io"
	"log"
	"math/big"
	"net/http"
	"time"
)

const balanceOfABI = `[{constant: true, inputs: [{name: "_owner", type: "address"}], name: "balanceOf", outputs: [{name: "balance", type: "uint256"}], payable: false, stateMutability: "view", type: "function"}]`

type Ethereum struct {
	redis  rueidis.Client
	config *Config
	esc    *etherscan.Client
	ec     *ethclient.Client
}

type UsdtTransaction struct {
	ID string
}

func New(cfg *Config, r rueidis.Client) (*Ethereum, error) {
	ethClient := etherscan.New(etherscan.Mainnet, cfg.ApiKey)
	nodeURL := cfg.ApiUrl + cfg.ApiKey
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		log.Fatal(err)
	}

	s := &Ethereum{
		redis:  r,
		config: cfg,
		esc:    ethClient,
		ec:     client,
	}

	return s, nil
}

func (e *Ethereum) GetUsdtContractEvents(lastScannedTrxID string) ([]TokenTransfer, error) {
	return e.getContractEventsByAddress(e.config.USDTAddress, lastScannedTrxID)
}

func (e *Ethereum) GetUsdcContractEvents(lastScannedTrxID string) ([]TokenTransfer, error) {
	return e.getContractEventsByAddress(e.config.USDCAddress, lastScannedTrxID)
}

func (e *Ethereum) getContractEventsByAddress(contract, lastScannedTrxID string) ([]TokenTransfer, error) {
	log.Printf("Start from %s\n", lastScannedTrxID)
	var allTokenTransfers []TokenTransfer
	page := 0

	for {
		tokenTransfers, err := e.fetchTokenTransfers(contract, page)
		if err != nil {
			log.Fatalf("Error fetching token transfers: %v", err)
		}

		visitedIndex := -1
		for index, transfer := range tokenTransfers {
			if transfer.Hash == lastScannedTrxID {
				visitedIndex = index
				break
			}
		}

		tokenTransfers = tokenTransfers[:visitedIndex]
		log.Printf("Number of transactions %d", len(tokenTransfers))
		allTokenTransfers = append(allTokenTransfers, tokenTransfers...)
		page++
		if len(tokenTransfers) < 1000 || visitedIndex != -1 {
			break
		}
		// Wait for the specified interval before the next request
		<-time.After(200 * time.Millisecond)
	}

	// Print all token transfers
	for i, transfer := range allTokenTransfers {
		fmt.Printf("%d. Transaction ID: %s, From: %s, To: %s, Amount: %s\n", i, transfer.Hash, transfer.From, transfer.To, transfer.Value)
	}

	return allTokenTransfers, nil
}
func (e *Ethereum) GetBalance(address, currency string) (*big.Float, error) {
	switch currency {
	case "ETH":
		return e.getEthBalance(address)
	case "USDT":
		return e.GetContractBalance(address, e.config.USDTAddress)
	case "USDC":
		return e.GetContractBalance(address, e.config.USDCAddress)
	default:
		return nil, fmt.Errorf("unsupported currency")
	}
}

func (e *Ethereum) getEthBalance(address string) (*big.Float, error) {
	walletAddress := common.HexToAddress(address)
	balance, err := e.ec.BalanceAt(context.Background(), walletAddress, nil)
	if err != nil {
		return new(big.Float), err
	}

	divisor := new(big.Float).SetPrec(128).SetFloat64(1e18)
	result := new(big.Float)
	result.Quo(new(big.Float).SetInt(balance), divisor)
	return result, nil
}

func (e *Ethereum) GetContractBalance(address, contract string) (*big.Float, error) {
	tokenAddress := common.HexToAddress(contract)
	instance, err := NewErc20(tokenAddress, e.ec)
	if err != nil {
		log.Fatal(err)
	}
	wallet := common.HexToAddress(address)
	balance, err := instance.BalanceOf(&bind.CallOpts{}, wallet)
	if err != nil {
		log.Fatal(err)
	}

	divisor := new(big.Float).SetPrec(128).SetFloat64(1e6)
	result := new(big.Float)
	result.Quo(new(big.Float).SetInt(balance), divisor)
	return result, nil
}

func (e *Ethereum) fetchTokenTransfers(contract string, page int) ([]TokenTransfer, error) {
	url := fmt.Sprintf("%s?module=account&action=tokentx&contractaddress=%s&sort=desc&apikey=%s&offset=1000",
		e.config.EtherscanApiUrl, contract, e.config.EtherscanApiKey)

	if page > 0 {
		url += fmt.Sprintf("&page=%d", page)
	}

	log.Printf("Fetched token transfers from %s", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Error closing response body: %v", err)
		}
	}(resp.Body)

	var responseBody io.Reader = resp.Body
	var body []byte
	body, err = io.ReadAll(responseBody)
	if err != nil {
		return nil, err
	}

	var tokenTransferResponse TokenTransferResponse
	err = json.Unmarshal(body, &tokenTransferResponse)
	if err != nil {
		return nil, err
	}

	return tokenTransferResponse.Result, nil
}

func (e *Ethereum) GetTransactionInfo(txId string) ([]map[string]interface{}, error) {

	return nil, nil
}

func (e *Ethereum) GetTrxBalance(address string) ([]map[string]interface{}, error) {

	return nil, nil
}

func (e *Ethereum) GetTransactions(address string) ([]map[string]interface{}, error) {

	return nil, nil
}

func (e *Ethereum) GetContractTransactions(address string) ([]map[string]interface{}, error) {

	return nil, nil
}

func (e *Ethereum) SendTrx(src Wallet, dest string, amount float64) ([]map[string]interface{}, error) {

	return nil, nil
}

func (e *Ethereum) SendUsdt(src Wallet, dest string, amount float64) ([]map[string]interface{}, error) {

	return nil, nil
}

func (e *Ethereum) createTransactionObject(log types.Log) (Transaction, error) {
	from := common.BytesToAddress(log.Topics[1].Bytes()).String()
	to := common.BytesToAddress(common.LeftPadBytes(log.Topics[2].Bytes(), 20)).String()
	amount := new(big.Int).SetBytes(log.Data).Int64()
	if amount == 0 {
		return Transaction{}, fmt.Errorf("amount is 0")
	}
	currency := ""
	switch log.Address.String() {
	case e.config.USDCAddress:
		currency = "USDC"
	case e.config.USDTAddress:
		currency = "USDT"
	default:
		return Transaction{}, fmt.Errorf("unknown contract address")
	}
	t := Transaction{
		TxID:            log.TxHash.String(),
		From:            from,
		To:              to,
		Amount:          amount,
		Currency:        currency,
		FeeLimit:        0,
		Timestamp:       0,
		ContractAddress: log.Address.String(),
	}

	return t, nil
}
