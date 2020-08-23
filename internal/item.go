package internal

import (
	"github.com/shopspring/decimal"
)

type Item struct {
	name     string
	price    decimal.Decimal
	selector int
}

func NewItem(n string, p float64, s int) Item {
	return Item{
		name:     n,
		price:    decimal.NewFromFloat(p),
		selector: s,
	}
}

type ItemRepository struct {
	Items []Item
}

func NewItemRepository() *ItemRepository {
	return &ItemRepository{
		Items: nil,
	}
}

func (r *ItemRepository) AddItem(i Item) {
	r.Items = append(r.Items, i)
}

func (r *ItemRepository) RemoveItem(item Item) {
	for i, rItem := range r.Items {
		if rItem.name == item.name {
			r.Items = append(r.Items[:i], r.Items[i+1:]...)
			break
		}
	}
}

func (r *ItemRepository) GetItemsAmount() int {
	return len(r.Items)
}
