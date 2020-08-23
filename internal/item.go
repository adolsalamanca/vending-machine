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

// Assumes price will be valid, thats why there is not error returned
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

func NewItemCatalogue() *ItemCatalogue {
	return &ItemCatalogue{
		Items: nil,
	}
}

func (r *ItemCatalogue) AddItem(i Item) {
	r.Items = append(r.Items, i)
}

func (r *ItemCatalogue) RemoveItem(item Item) {
	for i, rItem := range r.Items {
		if rItem.name == item.name {
			r.Items = append(r.Items[:i], r.Items[i+1:]...)
			break
		}
	}
}

func (r *ItemCatalogue) GetCatalogItemsAmount() int {
	return len(r.Items)
}

func (r *ItemCatalogue) GetItemByName(name string) (Item, error) {
	for _, i := range r.Items {
		if strings.EqualFold(name, i.name) {
			return i, nil
		}
	}

	return Item{}, NotFoundItemErr
}
