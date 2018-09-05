package main

type CoinJar interface {
	AddFunds(amount float64) error
}

type StarlingCoinJar struct {
	Name           string
	SavingsGoalUID string
}

func NewCoinJar(name string) CoinJar {
	savingsGoalUID, err := ensureStarlingSavingsGoal(name)
	if err != nil {
		panic("Unable to ensure a Starling Savings Goal exists")
	}

	return &StarlingCoinJar{
		Name:           name,
		SavingsGoalUID: savingsGoalUID,
	}
}

func (cj *StarlingCoinJar) AddFunds(amount float64) error {

	return nil
}

func ensureStarlingSavingsGoal(name string) (savingsGoalUID string, err error) {
	// get list of savings goals
	// check ours is in it
	// if not, make it

	return
}
