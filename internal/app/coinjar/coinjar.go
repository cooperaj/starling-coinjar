package coinjar

type CoinJar interface {
	AddFunds(amount int8) error
	GetRoundTo() int8
}
