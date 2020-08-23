package internal

import (
	"github.com/shopspring/decimal"
	"log"
	"strings"
)

const (
	Water = "Water"
	Juice = "Juice"
	Soda  = "Soda"

	WaterPrice = "0.65"
	JuicePrice = "1.00"
	SodaPrice  = "1.50"

	NotFoundItemErr = Error("Item not found")
)

type Item struct {
	selector int
	name     string
	price    decimal.Decimal
}

// NewItem instantiates a new item, it assumes price will be valid, so there is a log.Fatalf instead of error returned
func NewItem(n string, p string, s int) Item {
	price, err := decimal.NewFromString(p)
	if err != nil {
		log.Fatalf("Could not create new items")
	}
	return Item{
		selector: s,
		name:     n,
		price:    price,
	}
}

type ItemCatalogue struct {
	Items []Item
}

// NewItemCatalogue creates a Catalogue of items, which is an Array with some extended behavior
func NewItemCatalogue() *ItemCatalogue {
	return &ItemCatalogue{
		Items: nil,
	}
}

// AddItem inserts an item in the catalog
func (r *ItemCatalogue) AddItem(i Item) {
	r.Items = append(r.Items, i)
}

// RemoveItem removes an item from the catalog
func (r *ItemCatalogue) RemoveItem(item Item) {
	for i, rItem := range r.Items {
		if rItem.name == item.name {
			r.Items = append(r.Items[:i], r.Items[i+1:]...)
			break
		}
	}
}

// GetCatalogItemsAmount returns the total items in the current catalog
func (r *ItemCatalogue) GetCatalogItemsAmount() int {
	return len(r.Items)
}

// GetItemByName retrieves the Item instance that matches with the specified name ignoring case
func (r *ItemCatalogue) GetItemByName(name string) (Item, error) {
	for _, i := range r.Items {
		if strings.EqualFold(name, i.name) {
			return i, nil
		}
	}

	return Item{}, NotFoundItemErr
}
