package tron

import (
	"github.com/stedigate/core/pkg/blockchains"
	"math/big"
)

type Transaction struct {
	TxID            string
	BlockNumber     int64
	From            Wallet
	To              Wallet
	Amount          big.Float
	Symbol          blockchains.TokenSymbol
	FeeLimit        int
	Timestamp       int
	ContractAddress string
}
