package internal_test

import (
	"github.com/adolsalamanca/vending-machine/internal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	WATER = "Water"
	JUICE = "Juice"
	SODA  = "Soda"
)

var _ = Describe("Item test", func() {

	var item1, item2, item3 internal.Item
	var catalogue *internal.ItemCatalogue

	BeforeEach(func() {
		item1 = internal.NewItem(WATER, 0.65, 1)
		item2 = internal.NewItem(JUICE, 1.00, 2)
		item3 = internal.NewItem(SODA, 1.50, 3)

		catalogue = internal.NewItemCatalogue()

		catalogue.AddItem(item1)
		catalogue.AddItem(item2)
	})

	Context("Catalogue test", func() {

		When("add an item", func() {

			It("should contain also the added item", func() {
				catalogue.AddItem(item3)

				Expect(catalogue.GetCatalogItemsAmount()).To(BeEquivalentTo(3))
			})

		})

		When("remove an item", func() {

			It("should only contain not removed item", func() {
				catalogue.RemoveItem(item2)

				Expect(catalogue.GetCatalogItemsAmount()).To(BeEquivalentTo(1))
			})

		})

	})

	Context("Stock test", func() {

	})

})
