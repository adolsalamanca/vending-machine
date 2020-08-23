package internal_test

import (
	"github.com/adolsalamanca/vending-machine/internal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ItemRepository test", func() {

	var item1, item2, item3 internal.Item
	var repository *internal.ItemRepository

	BeforeEach(func() {
		item1 = internal.NewItem("Water", 0.65)
		item2 = internal.NewItem("Juice", 1.00)
		item3 = internal.NewItem("Soda", 1.50)

		repository = internal.NewItemRepository()

		repository.AddItem(item1)
		repository.AddItem(item2)
	})

	When("added some items", func() {

		It("should contain all the added items", func() {
			repository.AddItem(item3)

			Expect(repository.GetItemsNumber()).To(BeEquivalentTo(3))
		})

	})

	When("removed some items", func() {

		It("should contain all the added items", func() {
			repository.RemoveItem(item2)

			Expect(repository.GetItemsNumber()).To(BeEquivalentTo(1))
		})

	})

})
