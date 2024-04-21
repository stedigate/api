package blockchains

import (
	"math/big"
	"time"
)

type Transaction interface {
	GetTxId() (string, error)
	GetCurrency() (TokenSymbol, error)
	GetBlockchain() (Blockchain, error)
	GetStatus() (TransactionStatus, error)
	GetFromAddress() (string, error)
	GetToAddress() (string, error)
	GetAmount() (big.Float, error)
	GetCreatedAt() (time.Time, error)
	GetTransactionInfo() ([]map[string]interface{}, error)
}
