package tron

type Transaction struct {
	TxID            string
	From            Wallet
	To              Wallet
	Amount          float64
	Currency        string
	FeeLimit        int
	Timestamp       int
	ContractAddress string
}
