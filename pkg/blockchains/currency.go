package blockchains

type TokenSymbol string

const (
	ETH  TokenSymbol = "ETH"
	TRX  TokenSymbol = "TRX"
	SOL  TokenSymbol = "SOL"
	TON  TokenSymbol = "TON"
	AVAX TokenSymbol = "AVAX"
	USDT TokenSymbol = "USDT"
	EURT TokenSymbol = "EURT"
	USDC TokenSymbol = "USDC"
	EURC TokenSymbol = "EURC"
)

func NewTokenSymbol(currency string) TokenSymbol {
	c := TokenSymbol(currency)
	if !c.isValid() {
		panic("Invalid Token Symbol name " + currency)
	}

	return c
}

func (c TokenSymbol) String() string {
	return string(c)
}

func (c TokenSymbol) isValid() bool {
	switch c {
	case USDT, EURT, USDC, EURC, ETH, TRX, SOL, TON, AVAX:
		return true
	}
	return false
}

func Currencies() []TokenSymbol {
	return []TokenSymbol{USDT, EURT, USDC, EURC, ETH, TRX, SOL, TON, AVAX}
}
