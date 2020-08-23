package internal_test

import (
	"github.com/adolsalamanca/vending-machine/internal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"
)

const (
	storedValue = 1.40
)

var _ = Describe("coinEngine test", func() {

	var coin1, coin2, coin3, coin4 internal.Coin
	var engine *internal.CashEngine
	var err error

	BeforeEach(func() {
		coin1, err = internal.NewCoin(internal.FiveCent)
		Expect(err).To(BeNil())

		coin2, err = internal.NewCoin(internal.TenCent)
		Expect(err).To(BeNil())

		coin3, err = internal.NewCoin(internal.TwentyFiveCent)
		Expect(err).To(BeNil())

		coin4, err = internal.NewCoin(internal.OneUnit)
		Expect(err).To(BeNil())

		engine = internal.NewCashEngine()
		engine.StoreCoins(coin1, coin2, coin3, coin4)

	})

	When("insert and then drop some coins", func() {

		It("should return only coins inserted in last service", func() {
			engine.InsertCoins(coin4)

			coins := engine.DropCoins()

			Expect(coins).To(ContainElement(coin4))
			Expect(len(coins)).To(BeEquivalentTo(1))
		})

	})

	When("store some coins", func() {

		It("should increase balance inside the engine", func() {
			engine.StoreCoins(coin4)

			v, err := decimal.NewFromString(internal.OneUnit)
			Expect(err).To(BeNil())
			value, _ := v.Float64()

			currentStoredValue := storedValue + value
			Expect(engine.GetBalance().Float64()).To(BeEquivalentTo(currentStoredValue))
		})

	})

})
