package internal_test

import (
	"github.com/adolsalamanca/vending-machine/internal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Coin test", func() {

	When("created a coin with a valid value", func() {

		It("it should not return error", func() {
			_, err := internal.NewCoin(internal.FiveCent)

			Expect(err).To(BeNil())
		})
	})

	When("created a coin without a valid value", func() {

		It("it should return error", func() {
			invalidValue := "invalidValue"

			_, err := internal.NewCoin(invalidValue)

			Expect(err).ToNot(BeNil())
		})
	})

})
