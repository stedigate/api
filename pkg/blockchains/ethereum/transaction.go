package ethereum

type Transaction struct {
	TxID            string
	From            string
	To              string
	Amount          int64
	Currency        string
	Blockchain      string
	FeeLimit        int
	Timestamp       int
	ContractAddress string
}

type TokenTransfer struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	From              string `json:"from"`
	ContractAddress   string `json:"contractAddress"`
	To                string `json:"to"`
	Value             string `json:"value"`
	TokenName         string `json:"tokenName"`
	TokenSymbol       string `json:"tokenSymbol"`
	TokenDecimal      string `json:"tokenDecimal"`
	TransactionIndex  string `json:"transactionIndex"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	GasUsed           string `json:"gasUsed"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	Input             string `json:"input"`
	Confirmations     string `json:"confirmations"`
}
type TokenTransferResponse struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Result  []TokenTransfer `json:"result"`
}
