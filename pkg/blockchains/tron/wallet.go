package tron

import (
	"errors"
	"github.com/ranjbar-dev/tron-wallet/util"
)

type Wallet struct {
	privateKey    string
	publicKey     string
	Balance       float64
	AddressBase58 string
	AddressHex    string
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

func (w *Wallet) GetAddressBase58() string {
	if w.AddressBase58 == "" && w.AddressHex != "" {
		w.AddressBase58, _ = convertBase58ToHex(w.AddressHex)
	}

	return w.AddressBase58
}

func (w *Wallet) GetAddressHex() string {
	if w.AddressHex == "" && w.AddressBase58 != "" {
		w.AddressHex, _ = convertBase58ToHex(w.AddressBase58)
	}

	return w.AddressHex
}

func convertBase58ToHex(input string) (string, error) {
	addr, err := util.DecodeCheck(input)
	if err != nil {
		return "", err
	}
	return string(addr[:]), nil
}

func NewWallet(privateKey, base58 string) *Wallet {
	return &Wallet{
		privateKey:    privateKey,
		AddressBase58: base58,
	}
}

var (
	ErrDecodeLength = errors.New("base58 decode length error")
	ErrDecodeCheck  = errors.New("base58 check failed")
	ErrEncodeLength = errors.New("base58 encode length error")
)
