package coinjar_test

import (
	"fmt"

	. "github.com/cooperaj/starling-coinjar/internal/app/coinjar"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	Describe("Returning a correct minor unit value for a configuration string", func() {
		var configurationTestValues = []struct {
			configuredValue string
			decodedNumber   int8
		}{
			{"pound", 100},
			{"fifty", 50},
			{"twenty", 20},
			{"ten", 10},
			{"five", 5},
			{"hotgarbage", 100}, //tests default
		}

		for _, test := range configurationTestValues {
			Context(fmt.Sprintf("Given a configuration value of %s", test.configuredValue), func() {
				It(fmt.Sprintf("Should return a numeric value of %d", test.decodedNumber), func() {
					var config = RoundToDecoder(0)

					config.Decode(test.configuredValue)

					Expect(int8(config)).To(Equal(test.decodedNumber))
				})
			})
		}
	})
})
