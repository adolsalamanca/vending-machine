package internal

import (
	"github.com/shopspring/decimal"
	"sort"
)

const (
	NotEnoughCoinsErr   = Error("Not enough coins, please insert more")
	NotValidExchangeErr = Error("It is not possible to return exchange, please try with exact amount of coins")
)

type Error string

func (e Error) Error() string {
	return string(e)
}

type CashEngine struct {
	storedCashDetailed    map[string]int
	storedBalance         decimal.Decimal
	currentServiceCash    []Coin
	currentServiceBalance decimal.Decimal
	validCoins            []string
}

func NewCashEngine(validCoins ...string) *CashEngine {
	storedCashDetailed := make(map[string]int)
	for _, c := range validCoins {
		storedCashDetailed[c] = 0
	}

	sort.Sort(sort.Reverse(sort.StringSlice(validCoins)))

	return &CashEngine{
		storedCashDetailed:    storedCashDetailed,
		storedBalance:         decimal.Decimal{},
		currentServiceCash:    nil,
		currentServiceBalance: decimal.Decimal{},
		validCoins:            validCoins,
	}
}

func (e *CashEngine) InsertCoins(coins ...Coin) {
	for _, c := range coins {
		e.currentServiceCash = append(e.currentServiceCash, c)
		e.currentServiceBalance = e.currentServiceBalance.Add(c.value)
	}
}

func (e *CashEngine) DropCoins() []Coin {
	serviceCoins := e.currentServiceCash
	e.currentServiceCash = nil
	e.currentServiceBalance = decimal.Decimal{}

	return serviceCoins
}

func (e *CashEngine) StoreCoins(coins ...Coin) {
	for _, c := range coins {
		e.storedCashDetailed[c.category] += 1
		e.storedBalance = e.storedBalance.Add(c.value)
	}
}

func (e *CashEngine) GetBalance() decimal.Decimal {
	return e.storedBalance
}

func (e *CashEngine) GetPaid(price decimal.Decimal) ([]Coin, error) {
	if price.GreaterThan(e.currentServiceBalance) {
		return []Coin{}, NotEnoughCoinsErr
	}

	difference := e.currentServiceBalance.Sub(price)

	return e.GiveExchange(difference)
}

func (e *CashEngine) GiveExchange(difference decimal.Decimal) ([]Coin, error) {
	var exchange []Coin
	for _, currentCoin := range e.validCoins {
		validCoinValue, err := decimal.NewFromString(currentCoin)
		if err != nil {
			return []Coin{}, err
		}

		if validCoinValue.GreaterThan(difference) {
			continue
		}

		coinReturned, err := NewCoin(currentCoin)
		if err != nil {
			return []Coin{}, err
		}

		if difference.Mod(validCoinValue).Equal(decimal.Zero) {
			neededCoins := int(difference.Div(validCoinValue).IntPart())
			if e.storedCashDetailed[currentCoin] >= neededCoins {
				for i := 0; i < neededCoins; i++ {
					exchange = append(exchange, coinReturned)
					e.storedCashDetailed[currentCoin]--
				}

				return exchange, nil
			}

			difference = difference.Sub(validCoinValue.Mul(decimal.NewFromInt(int64(e.storedCashDetailed[currentCoin]))))
			for i := 0; i < e.storedCashDetailed[currentCoin]; i++ {
				exchange = append(exchange, coinReturned)
				e.storedCashDetailed[currentCoin]--
			}

		} else {
			var usedCoins int
			neededCoins := int(difference.Div(validCoinValue).IntPart())
			if e.storedCashDetailed[currentCoin] >= neededCoins {
				usedCoins = neededCoins
			} else {
				usedCoins = e.storedCashDetailed[currentCoin]
			}

			difference = difference.Sub(validCoinValue.Mul(decimal.NewFromInt(int64(usedCoins))))
			for i := 0; i < usedCoins; i++ {
				exchange = append(exchange, coinReturned)
				e.storedCashDetailed[currentCoin]--
			}
		}

	}

	if difference.Equal(decimal.Zero) {
		return exchange, nil
	}

	return []Coin{}, NotValidExchangeErr
}
