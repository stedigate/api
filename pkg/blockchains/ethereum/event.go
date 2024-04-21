package ethereum

type TransferEvent struct {
	BlockNumber     int    `json:"block_number"`
	BlockTimestamp  int64  `json:"block_timestamp"`
	ContractAddress string `json:"contract_address"`
	From            string `json:"from"`
	To              string `json:"to"`
	Amount          int64  `json:"amount"`
	TransactionID   string `json:"transaction_id"`
}
