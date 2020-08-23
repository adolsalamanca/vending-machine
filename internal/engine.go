package internal

import (
	"github.com/shopspring/decimal"
)

type CashEngine struct {
	storedCash         []Coin
	currentServiceCash []Coin
	balance            decimal.Decimal
}

func NewCashEngine() *CashEngine {
	return &CashEngine{
		storedCash:         nil,
		currentServiceCash: nil,
		balance:            decimal.Decimal{},
	}
}

func (e *CashEngine) InsertCoins(coins ...Coin) {
	for _, c := range coins {
		e.currentServiceCash = append(e.currentServiceCash, c)
	}
}

func (e *CashEngine) DropCoins() []Coin {
	serviceCoins := e.currentServiceCash
	e.currentServiceCash = nil

	return serviceCoins
}

func (e *CashEngine) StoreCoins(coins ...Coin) {
	for _, c := range coins {
		e.storedCash = append(e.storedCash, c)
		e.balance = e.balance.Add(c.value)
	}
}

func (e *CashEngine) GetBalance() decimal.Decimal {
	return e.balance
}
