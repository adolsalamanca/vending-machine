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

var _ = Describe("CashEngine test", func() {

	var fiveCentCoin, tenCentCoin, twentyFiveCentCoin, oneUnitCoin internal.Coin
	var engine *internal.CashEngine
	var err error

	BeforeEach(func() {
		fiveCentCoin, err = internal.NewCoin(internal.FiveCent)
		Expect(err).To(BeNil())

		tenCentCoin, err = internal.NewCoin(internal.TenCent)
		Expect(err).To(BeNil())

		twentyFiveCentCoin, err = internal.NewCoin(internal.TwentyFiveCent)
		Expect(err).To(BeNil())

		oneUnitCoin, err = internal.NewCoin(internal.OneUnit)
		Expect(err).To(BeNil())

		engine = internal.NewCashEngine()
		engine.StoreCoins(fiveCentCoin, tenCentCoin, twentyFiveCentCoin, oneUnitCoin)

	})

	When("insert and then drop some coins", func() {

		It("should return only coins inserted in last service", func() {
			engine.InsertCoins(oneUnitCoin)

			coins := engine.DropCoins()

			Expect(coins).To(ContainElement(oneUnitCoin))
			Expect(len(coins)).To(BeEquivalentTo(1))
		})

	})

	When("store some coins", func() {

		It("should increase storedBalance inside the engine", func() {
			engine.StoreCoins(oneUnitCoin)

			v, err := decimal.NewFromString(internal.OneUnit)
			Expect(err).To(BeNil())
			value, _ := v.Float64()

			currentStoredValue := storedValue + value
			balance, err := engine.GetBalance()

			Expect(err).To(BeNil())
			b, _ := balance.Float64()

			Expect(b).To(BeEquivalentTo(currentStoredValue))
		})

	})

	When("getting paid", func() {
		It("should return err if not enough money was inserted", func() {
			sodaPrice := decimal.NewFromFloat(1.50)
			engine.InsertCoins(oneUnitCoin)

			_, err := engine.GetPaid(sodaPrice)

			Expect(err).To(BeEquivalentTo(internal.NotEnoughCoinsErr))
		})

		It("should return err if all allowed coins are bigger than exchange difference", func() {
			engine = internal.NewCashEngine(internal.OneUnit)

			_, err := engine.GiveExchange(decimal.NewFromFloat(0.50))

			Expect(err).To(BeEquivalentTo(internal.NotValidExchangeErr))
		})

		It("should return NotValidExchangeErr err as it is not possible to return exchange due to lack of coins", func() {
			engine = internal.NewCashEngine(internal.OneUnit, internal.TenCent, internal.TwentyFiveCent, internal.FiveCent)
			engine.StoreCoins(twentyFiveCentCoin, tenCentCoin, fiveCentCoin, fiveCentCoin)

			sodaPrice := decimal.NewFromFloat(1.50)
			engine.InsertCoins(oneUnitCoin, oneUnitCoin)

			_, err := engine.GetPaid(sodaPrice)

			Expect(err).To(BeEquivalentTo(internal.NotValidExchangeErr))
		})

		It("should return exchange as it has enough coins of same type to give return", func() {
			engine = internal.NewCashEngine(internal.OneUnit, internal.TenCent, internal.TwentyFiveCent, internal.FiveCent)
			engine.StoreCoins(twentyFiveCentCoin, twentyFiveCentCoin)

			sodaPrice := decimal.NewFromFloat(1.50)
			engine.InsertCoins(oneUnitCoin, oneUnitCoin)

			coins, err := engine.GetPaid(sodaPrice)

			Expect(err).To(BeNil())
			Expect(len(coins)).To(BeEquivalentTo(2))
			Expect(coins).To(ContainElement(twentyFiveCentCoin))
		})

		It("should return exchange but of different coin types as it is what it has", func() {
			engine = internal.NewCashEngine(internal.OneUnit, internal.TenCent, internal.TwentyFiveCent, internal.FiveCent)
			engine.StoreCoins(twentyFiveCentCoin, tenCentCoin, tenCentCoin, fiveCentCoin)

			sodaPrice := decimal.NewFromFloat(1.50)
			engine.InsertCoins(oneUnitCoin, oneUnitCoin)

			coins, err := engine.GetPaid(sodaPrice)

			Expect(err).To(BeNil())
			Expect(len(coins)).To(BeEquivalentTo(4))
		})

		It("should return exchange but of different coin types as it is what it has", func() {
			engine = internal.NewCashEngine(internal.OneUnit, internal.TenCent, internal.TwentyFiveCent, internal.FiveCent)
			engine.StoreCoins(twentyFiveCentCoin, fiveCentCoin, fiveCentCoin, fiveCentCoin, fiveCentCoin, fiveCentCoin, fiveCentCoin)

			sodaPrice := decimal.NewFromFloat(1.50)
			engine.InsertCoins(oneUnitCoin, oneUnitCoin)

			coins, err := engine.GetPaid(sodaPrice)

			Expect(err).To(BeNil())
			Expect(len(coins)).To(BeEquivalentTo(6))
		})

		It("should return exchange using the biggest coins possible", func() {
			engine = internal.NewCashEngine(internal.OneUnit, internal.TenCent, internal.TwentyFiveCent, internal.FiveCent)
			engine.StoreCoins(twentyFiveCentCoin, tenCentCoin, fiveCentCoin, fiveCentCoin, fiveCentCoin)

			sodaPrice := decimal.NewFromFloat(1.50)
			engine.InsertCoins(oneUnitCoin, oneUnitCoin)

			coins, err := engine.GetPaid(sodaPrice)

			Expect(err).To(BeNil())
			Expect(len(coins)).To(BeEquivalentTo(5))
		})

	})

})
