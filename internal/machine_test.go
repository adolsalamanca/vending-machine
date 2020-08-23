package internal_test

import (
	"bytes"
	"github.com/adolsalamanca/vending-machine/internal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
)

var _ = Describe("Machine test", func() {

	var water, juice, soda internal.Item
	var catalogue *internal.ItemCatalogue
	var engine *internal.CashEngine
	var machine *internal.Machine
	var buf bytes.Buffer

	BeforeEach(func() {
		water = internal.NewItem(internal.Water, internal.WaterPrice, 1)
		juice = internal.NewItem(internal.Juice, internal.JuicePrice, 2)
		soda = internal.NewItem(internal.Soda, internal.SodaPrice, 3)

		catalogue = internal.NewItemCatalogue()

		catalogue.AddItem(water)
		catalogue.AddItem(juice)
		catalogue.AddItem(soda)

		engine = internal.NewCashEngine(internal.FiveCent, internal.TenCent, internal.TwentyFiveCent, internal.OneUnit)

		logger := log.Logger{}
		logger.SetOutput(&buf)

		machine = internal.NewMachine(catalogue, engine, &logger)
	})

	AfterEach(func() {
		buf.Reset()
	})

	Context("add coins & drop", func() {

		It("should return error if an illegal coin amount is inserted", func() {

			err := machine.InsertCoins(internal.TwoUnit)

			Expect(err).To(BeEquivalentTo(internal.NotValidCoinAmount))

		})

		It("should return error if an illegal coin type is inserted", func() {

			err := machine.InsertCoins("InvalidCoin")

			Expect(err).To(BeEquivalentTo(internal.NotValidCoin))

		})

		It("should drop the same amount of coins previously inserted", func() {
			expectedPrint := `-> 0.05, 0.1
`
			err := machine.InsertCoins(internal.FiveCent, internal.TenCent)

			Expect(err).To(BeNil())
			machine.ReturnCoins()

			Expect(buf.String()).To(BeEquivalentTo(expectedPrint))
		})

	})

	Context("sell item", func() {

		BeforeEach(func() {
			err := machine.LoadMoney(internal.FiveCent, internal.TenCent, internal.TwentyFiveCent, internal.OneUnit)
			Expect(err).To(BeNil())
			err = machine.LoadMoney(internal.FiveCent, internal.TenCent, internal.TwentyFiveCent, internal.OneUnit)
			Expect(err).To(BeNil())
			err = machine.LoadMoney(internal.FiveCent, internal.TenCent, internal.TwentyFiveCent, internal.OneUnit)
			Expect(err).To(BeNil())
			err = machine.LoadMoney(internal.FiveCent, internal.TenCent, internal.TwentyFiveCent, internal.OneUnit)
			Expect(err).To(BeNil())
			err = machine.LoadMoney(internal.FiveCent, internal.TenCent, internal.TwentyFiveCent, internal.OneUnit)
			Expect(err).To(BeNil())

			machine.LoadItem(water, 5)
			machine.LoadItem(juice, 5)
			machine.LoadItem(soda, 5)
		})

		FIt("reproduce issue", func() {
			err := machine.InsertCoins(internal.OneUnit, internal.TwentyFiveCent, internal.TwentyFiveCent)
			Expect(err).To(BeNil())
			err = machine.SellItem(internal.Soda)

			buf.Reset()

			err = machine.InsertCoins(internal.OneUnit, internal.TwentyFiveCent, internal.TwentyFiveCent)
			Expect(err).To(BeNil())
			err = machine.SellItem(internal.Soda)

		})

		It("should print expected response after non exact change sold item", func() {
			expectedPrint := `-> WATER, 0.25, 0.1
`
			err := machine.InsertCoins(internal.OneUnit)
			Expect(err).To(BeNil())

			err = machine.SellItem(internal.Water)

			Expect(err).To(BeNil())
			Expect(buf.String()).To(BeEquivalentTo(expectedPrint))

		})

		It("should print expected response after exact change sold item", func() {
			expectedPrint := `-> SODA
`
			err := machine.InsertCoins(internal.OneUnit, internal.TwentyFiveCent, internal.TwentyFiveCent)
			Expect(err).To(BeNil())

			err = machine.SellItem(internal.Soda)

			Expect(err).To(BeNil())
			Expect(buf.String()).To(BeEquivalentTo(expectedPrint))
		})

	})

})
