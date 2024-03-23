package tron

import (
	"encoding/json"
	"fmt"
	"github.com/redis/rueidis"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Tron struct {
	redis  rueidis.Client
	config *Config
}

type UsdtTransaction struct {
	ID string
}

func New(cfg *Config, r rueidis.Client) (*Tron, error) {
	t := &Tron{
		redis:  r,
		config: cfg,
	}

	return t, nil
}

func (t *Tron) GetContractEvents(lastScannedTrxID string) ([]TransferEvent, error) {
	url := "/v1/contracts/" + t.config.Trc20ContractAddress + "/events"

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

			fmt.Println(te.TransactionID)
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

func (t *Tron) GetTrxBalance(address string) ([]map[string]interface{}, error) {

	return nil, nil
}

func (t *Tron) GetContractBalance(address string) ([]map[string]interface{}, error) {

	return nil, nil
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

func (t *Tron) SendTrx(src Wallet, dest string, amount float64) ([]map[string]interface{}, error) {

	return nil, nil
}

func (t *Tron) SendUsdt(src Wallet, dest string, amount float64) ([]map[string]interface{}, error) {

	return nil, nil
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

		fmt.Println(req.URL.String())
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
