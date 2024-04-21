package blockchains

type Blockchain string

const (
	ETHEREUM  Blockchain = "Ethereum"
	AVALANCHE Blockchain = "Avalanche"
	SOLANA    Blockchain = "Solana"
	TRON      Blockchain = "Tron"
)

func NewBlockchain(n string) Blockchain {
	b := Blockchain(n)
	if !b.isValid() {
		panic("Invalid network name " + n)
	}

	return Blockchain(n)
}

func (b Blockchain) String() string {
	return string(b)
}

func (b Blockchain) isValid() bool {
	switch b {
	case ETHEREUM, AVALANCHE, SOLANA, TRON:
		return true
	}
	return false
}

func Blockchains() []Blockchain {
	return []Blockchain{ETHEREUM, AVALANCHE, SOLANA, TRON}
}
