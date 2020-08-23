package internal

import (
	"github.com/shopspring/decimal"
	"strings"
)

const NotFoundItemErr = Error("Item not found")

type Item struct {
	selector int
	name     string
	price    decimal.Decimal
}

func NewItem(n string, p float64, s int) Item {
	return Item{
		selector: s,
		name:     n,
		price:    decimal.NewFromFloat(p),
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
