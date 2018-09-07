package main

import "math"

// The different values we can calculate rounding change from
const (
	ChangeToAPound      = 100
	ChangeToFiftyPence  = 50
	ChangeToTwentyPence = 20
	ChangeToTenPence    = 10
	ChangeToFivePence   = 5
)

// CalculateChange works out how many pennies we should put in a coin jar
// from a given transaction
func CalculateChange(transaction Transaction, roundTo int8) int8 {
	var change = 0

	var amount = convertFloatToMinorUnits(transaction.Amount)
	var roundValue = int(roundTo)
	var remainder = amount % roundValue

	if remainder > 0 {
		change = roundValue - remainder
	}

	return int8(change)
}

func convertFloatToMinorUnits(amount float64) int {
	return int(math.Round(amount * -100))
}
