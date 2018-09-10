package coinjar_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestStarlingCoinjar(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "StarlingCoinjar Suite")
}
