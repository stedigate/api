package blockchains

type TransactionStatus string

const (
	Confirming TransactionStatus = "Confirming"
	Failed     TransactionStatus = "Failed"
	Confirmed  TransactionStatus = "Confirmed"
)

func NewTransactionStatus(status string) TransactionStatus {
	ts := TransactionStatus(status)
	if !ts.isValid() {
		panic("Invalid transaction status " + status)
	}

	return ts
}

func (ts TransactionStatus) String() string {
	return string(ts)
}

func (ts TransactionStatus) isValid() bool {
	switch ts {
	case Confirming, Failed, Confirmed:
		return true
	}
	return false
}

func TransactionStatuses() []TransactionStatus {
	return []TransactionStatus{Confirming, Failed, Confirmed}
}
