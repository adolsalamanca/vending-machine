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

}

func (r *ItemRepository) RemoveItem(i Item) {

}

func (r *ItemRepository) GetItemsNumber() int {
	return len(r.Items)
}
