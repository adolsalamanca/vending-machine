package internal

import (
	"github.com/shopspring/decimal"
)

type Item struct {
	Name  string
	Price decimal.Decimal
}

func NewItem(n string, p float64) Item {
	return Item{
		Name:  n,
		Price: decimal.NewFromFloat(p),
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
		if rItem.Name == item.Name {
			r.Items = append(r.Items[:i], r.Items[i+1:]...)
			break
		}
	}
}

func (r *ItemRepository) GetItemsAmount() int {
	return len(r.Items)
}
