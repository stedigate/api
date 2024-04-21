package dto

type Transaction struct {
	TxID            string
	From            string
	To              string
	Amount          float64
	Currency        string
	Blockchain      string
	FeeLimit        int
	Timestamp       int
	ContractAddress string
}
