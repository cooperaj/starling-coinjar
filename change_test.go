package main_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cooperaj/starling-coinjar"
)

var _ = Describe("Change", func() {
	Describe("Returning a correct amount of coin jar change given a transaction value", func() {
		var changeTestValues = []struct {
			transaction float64
			rounding    int8
			change      int8
		}{
			{-24.99, ChangeToAPound, 1},
			{-24.99, ChangeToFiftyPence, 1},
			{-24.99, ChangeToTwentyPence, 1},
			{-24.99, ChangeToTenPence, 1},
			{-24.99, ChangeToFivePence, 1},
			{-1.00, ChangeToAPound, 0},
			{-1.00, ChangeToFiftyPence, 0},
			{-1.00, ChangeToTwentyPence, 0},
			{-1.00, ChangeToTenPence, 0},
			{-1.00, ChangeToFivePence, 0},
			{-0.99, ChangeToAPound, 1},
			{-0.99, ChangeToFiftyPence, 1},
			{-0.99, ChangeToTwentyPence, 1},
			{-0.99, ChangeToTenPence, 1},
			{-0.99, ChangeToFivePence, 1},
			{-0.85, ChangeToAPound, 15},
			{-0.85, ChangeToFiftyPence, 15},
			{-0.85, ChangeToTwentyPence, 15},
			{-0.85, ChangeToTenPence, 5},
			{-0.85, ChangeToFivePence, 0},
			{-0.02, ChangeToAPound, 98},
			{-0.02, ChangeToFiftyPence, 48},
			{-0.02, ChangeToTwentyPence, 18},
			{-0.02, ChangeToTenPence, 8},
			{-0.02, ChangeToFivePence, 3},
		}

		for _, test := range changeTestValues {
			Context(fmt.Sprintf("Given a transaction value of %f and rounding to %d", test.transaction, test.rounding), func() {
				It(fmt.Sprintf("Should return a change value of %d", test.change), func() {
					var transaction = Transaction{
						Amount: float64(test.transaction),
						Type:   "TEST_TRANSACTION",
					}

					var change = CalculateChange(transaction, test.rounding)
					Expect(change).To(Equal(test.change))
				})
			})
		}
	})
})
