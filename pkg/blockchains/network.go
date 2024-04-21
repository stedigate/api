package blockchains

import "math/big"

type Network interface {
	Track(token TokenSymbol) ([]string, error)
	GetTransactionById(txId string) (Transaction, error)
	GetTransactionsByWallet(wallet Wallet) (Transaction, error)
	GetTransactionsByBlock(wallet Wallet) (Transaction, error)
	GetBalance(address string, token TokenSymbol) (TokenSymbol, error)
	GetTransactionStatus(txId string) (TransactionStatus, error)
	Send(from, to string, amount big.Float, token TokenSymbol) (Transaction, error)
}
