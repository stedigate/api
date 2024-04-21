package blockchains

import "math/big"

type Wallet interface {
	Balance(c TokenSymbol) big.Float
	Transactions() []Transaction
	Generate() string
	GetAddress() string
	GetPrivateKey() string
	GetPublicKey() string
}
