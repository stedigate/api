package tron

import (
	"encoding/json"
	"fmt"
	tronWallet "github.com/ranjbar-dev/tron-wallet"
	"github.com/ranjbar-dev/tron-wallet/enums"
	"github.com/redis/rueidis"
	"io"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const getEventsByContract = "/v1/contracts/%s/events"
const getBalanceURL = "/wallet/getaccount"
const GET_BALANCE_URL = "/v1/accounts/%s/balances"
const CREATED_TRANSACTION = "/wallet/createtransaction"

var usdtToken *tronWallet.Token

type Tron struct {
	redis  rueidis.Client
	config *Config
}

type UsdtTransaction struct {
	ID string
}

func New(cfg *Config, r rueidis.Client) (*Tron, error) {
	err := os.Setenv("TRON_PRO_API_KEY", cfg.TrongridApiKey)
	if err != nil {
		return nil, err
	}
	usdtToken = &tronWallet.Token{
		ContractAddress: enums.ContractAddress(cfg.USDTAddress),
	}

	t := &Tron{
		redis:  r,
		config: cfg,
	}

	return t, nil
}

func (t *Tron) GetContractEvents(lastScannedTrxID string) ([]TransferEvent, error) {
	url := fmt.Sprintf(getEventsByContract, t.config.USDTAddress)
	body := map[string]interface{}{
		"only_confirmed": true,
		"limit":          200,
		"order_by":       "block_timestamp,desc",
		"event_name":     "Transfer",
	}

	var events []TransferEvent

	for {
		res, err := t.sendRequest(url, body, "GET")
		if err != nil {
			return events, fmt.Errorf("unable to send request: %w", err)
		}

		var data map[string]interface{}
		err = json.Unmarshal(res, &data)
		if err != nil {
			return events, fmt.Errorf("unable to unmarshal response: %w", err)
		}

		if data["success"] != true {
			return events, fmt.Errorf("request failed: %v", data)
		}

		for _, ev := range data["data"].([]interface{}) {
			amount, _ := strconv.ParseInt(
				ev.(map[string]interface{})["result"].(map[string]interface{})["value"].(string),
				10,
				64,
			)
			te := TransferEvent{
				BlockNumber:     int(ev.(map[string]interface{})["block_number"].(float64)),
				BlockTimestamp:  int64(ev.(map[string]interface{})["block_timestamp"].(float64)),
				ContractAddress: ev.(map[string]interface{})["contract_address"].(string),
				From:            ev.(map[string]interface{})["result"].(map[string]interface{})["from"].(string),
				To:              ev.(map[string]interface{})["result"].(map[string]interface{})["to"].(string),
				Amount:          amount,
				TransactionID:   ev.(map[string]interface{})["transaction_id"].(string),
			}
			if te.TransactionID == lastScannedTrxID {
				return events, nil
			}
			events = append(events, te)
		}

		fingerprint, ok := data["meta"].(map[string]interface{})["fingerprint"]
		if ok {
			body["fingerprint"] = fingerprint
		} else {
			return events, nil
		}
	}
}

func (t *Tron) GetTransactionInfo(txId string) ([]map[string]interface{}, error) {

	return nil, nil
}

func (t *Tron) GetBalance(address, currency string) (*big.Float, error) {
	w, err := t.createWalletFromAddress(address)
	if err != nil {
		return new(big.Float), err
	}

	switch currency {
	case "TRX":

		balance, err := w.Balance()
		if err != nil {
			return new(big.Float), err
		}
		divisor := new(big.Float).SetPrec(128).SetFloat64(1e6)
		result := new(big.Float)
		result.Quo(new(big.Float).SetInt64(balance), divisor)
		return result, nil
	case "USDT":
		balance, err := w.BalanceTRC20(usdtToken)
		if err != nil {
			return new(big.Float), err
		}
		divisor := new(big.Float).SetPrec(128).SetFloat64(1e6)
		result := new(big.Float)
		result.Quo(new(big.Float).SetInt64(balance), divisor)
		return result, nil
	default:
		return nil, fmt.Errorf("unsupported currency")
	}
}

func (t *Tron) createWalletFromAddress(address string) (*tronWallet.TronWallet, error) {
	node := enums.Node(t.config.TrongridGrpcUrl)
	var w *tronWallet.TronWallet
	var err error
	if len(address) == 34 {
		w = &tronWallet.TronWallet{
			Node:          node,
			Address:       "",
			AddressBase58: address,
			PrivateKey:    "",
			PublicKey:     address,
		}
	} else {
		w, err = tronWallet.CreateTronWallet(node, address)
	}

	return w, err
}

func (t *Tron) GetTransactions(address string) ([]map[string]interface{}, error) {

	return nil, nil
}

func (t *Tron) GetContractTransactions(address string) ([]map[string]interface{}, error) {

	return nil, nil
}

func (t *Tron) ValidateAddress(address string) (bool, error) {
	url := "/wallet/validateaddress"
	body := map[string]interface{}{
		"address": address,
		"visible": true,
	}
	response, err := t.sendRequest(url, body, "POST")
	if err != nil {
		return false, fmt.Errorf("unable to send request: %w", err)
	}

	var data map[string]interface{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return false, fmt.Errorf("unable to send request: %w", err)
	}

	return data["result"].(bool), nil
}

func (t *Tron) Send(src, dest, currency string, amount uint64) (string, error) {
	switch currency {
	case "TRX":
		return t.sendTrx(src, dest, amount)
	case "USDT":
		return t.sendTrc20(src, dest, amount)
	default:
		return "", fmt.Errorf("unsupported currency")
	}
}

func (t *Tron) sendTrx(src, dest string, a uint64) (string, error) {
	w, err := t.createWalletFromAddress(src)
	if err != nil {
		return "", err
	}

	txId, err := w.Transfer(dest, int64(a*1e6))
	if err != nil {
		return "", err
	}

	return txId, nil
}

func (t *Tron) sendTrc20(src, dest string, a uint64) (string, error) {
	w, err := t.createWalletFromAddress(src)
	if err != nil {
		return "", err
	}

	txId, err := w.TransferTRC20(usdtToken, dest, int64(a*1e6))
	if err != nil {
		return "", err
	}

	return txId, nil
}

func (t *Tron) SimulateSend(src, dest, currency string, amount uint64) (*big.Float, error) {
	w, err := t.createWalletFromAddress(src)
	if err != nil {
		return new(big.Float), err
	}

	switch currency {
	case "TRX":
		feeInSun, err := w.EstimateTransferFee(dest, int64(amount*1e6))
		if err != nil {
			return new(big.Float), err
		}
		divisor := new(big.Float).SetPrec(128).SetFloat64(1e6)
		result := new(big.Float)
		result.Quo(new(big.Float).SetInt64(feeInSun), divisor)
		return result, nil
	case "USDT":
		feeInSun, err := w.EstimateTransferTRC20Fee()
		if err != nil {
			return new(big.Float), err
		}
		divisor := new(big.Float).SetPrec(128).SetFloat64(1e6)
		result := new(big.Float)
		result.Quo(new(big.Float).SetInt64(feeInSun), divisor)
		return result, nil
	default:
		return new(big.Float), fmt.Errorf("unsupported currency")
	}
}

func (t *Tron) simulateSendTrx(src, dest string, a uint64) (string, error) {
	w, err := t.createWalletFromAddress(src)
	if err != nil {
		return "", err
	}

	txId, err := w.Transfer(dest, int64(a*1e6))
	if err != nil {
		return "", err
	}

	return txId, nil
}

func (t *Tron) simulateSendTrc20(src, dest string, a uint64) (string, error) {
	w, err := t.createWalletFromAddress(src)
	if err != nil {
		return "", err
	}

	txId, err := w.TransferTRC20(usdtToken, dest, int64(a*1e6))
	if err != nil {
		return "", err
	}

	return txId, nil
}

func (t *Tron) generateJwtToken() (string, error) {
	/*dataBytes, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("unable to marshal data: %w", err)
	}

	expiredAt := jwt.NewNumericDate(time.Now().Add(t.config.JwtExpiration))
	registeredClaim := jwt.RegisteredClaims{
		ExpiresAt: expiredAt,
		Audience:  jwt.ClaimStrings{"trongrid.io"},
	}
	claims := &JwtClaims{dataBytes, registeredClaim}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	jwtToken.Header["kid"] = t.config.TrongridJwtKeyId
	jwtToken.Header["typ"] = "JWT"

	token, err := jwtToken.SignedString(t.config.JwtPrivateKey)
	if err != nil {
		return "", fmt.Errorf("unable to marshal data: %w", err)
	}*/

	token := t.config.TrongridJwtToken
	return token, nil
}

func (t *Tron) sendRequest(path string, body map[string]interface{}, method string) ([]byte, error) {
	token, _ := t.generateJwtToken()

	req, _ := http.NewRequest(method, t.config.TrongridApiUrl+path, nil)

	// Set the headers
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "com.stedigate.app/v1.0.0 (by Liandm Ltd)")
	req.Header.Add("TRON-PRO-API-KEY", t.config.TrongridApiKey)
	req.Header.Add("Origin", "https://www.stedigate.io")

	// Set the query parameters
	if method == "GET" {
		q := req.URL.Query()
		for key, value := range body {
			q.Add(key, fmt.Sprint(value))
		}
		req.URL.RawQuery = q.Encode()
	} else {
		payload, _ := json.Marshal(body)
		req.Body = io.NopCloser(strings.NewReader(string(payload)))
	}

	res, _ := http.DefaultClient.Do(req)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	return io.ReadAll(res.Body)
}
