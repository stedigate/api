package solana

type Wallet struct {
	privateKey    string
	publicKey     string
	Balance       float64
	AddressBase58 string
}

func (w *Wallet) GetPrivateKey() string {
	return w.privateKey
}

func (w *Wallet) GetPublicKey() string {
	return w.publicKey
}

func (w *Wallet) GetBalance() float64 {
	return w.Balance
}

func (w *Wallet) GetAddress() string {
	return w.AddressBase58
}

func convertBase58ToHex(base58 string) string {
	return ""
}

func NewWallet(privateKey, base58 string) *Wallet {
	return &Wallet{
		privateKey:    privateKey,
		AddressBase58: base58,
	}
}
