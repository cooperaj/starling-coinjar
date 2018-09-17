package coinjar

type Transaction struct {
	Amount float64
	Type   string
}

type TransactionProcessor interface {
	Start()
	Stop()
	ProcessPayload(payload []byte) error
}
