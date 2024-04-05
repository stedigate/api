package solana

type Direction int32

const (
	Unknown Direction = iota
	From    Direction = iota
	To
	Both
)

func (d Direction) String() string {
	switch d {
	case From:
		return "from"
	case To:
		return "to"
	case Both:
		return "both"
	default:
		return "unknown"
	}
}
