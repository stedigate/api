package solana

import (
	"github.com/stedigate/core/pkg/blockchains"
	"math/big"
	"time"
)

type Transaction struct {
	TxID            string
	BlockNumber     uint64
	From            Wallet
	To              Wallet
	Amount          big.Float
	Blockchain      blockchains.Blockchain
	Status          blockchains.TransactionStatus
	Timestamp       time.Time
	FeeLimit        uint64
	ContractAddress string
	Symbol          blockchains.TokenSymbol
}

func (t *Transaction) GetTxID() (string, error) {
	return t.TxID, nil
}

func (t *Transaction) GetCurrency() (blockchains.TokenSymbol, error) {
	return t.Symbol, nil
}

func (t *Transaction) GetBlockchain() (blockchains.Blockchain, error) {
	return t.Blockchain, nil
}

func (t *Transaction) GetStatus() (blockchains.TransactionStatus, error) {
	return t.Status, nil
}

func (t *Transaction) GetFromAddress() (string, error) {
	return t.From.GetAddress(), nil
}

func (t *Transaction) GetToAddress() (string, error) {
	return t.To.GetAddress(), nil
}

func (t *Transaction) GetAmount() (big.Float, error) {
	return t.Amount, nil
}

func (t *Transaction) GetCreatedAt() (time.Time, error) {
	return t.Timestamp, nil
}

func (t *Transaction) GetTransactionInfo() ([]map[string]interface{}, error) {
	return nil, nil
}
