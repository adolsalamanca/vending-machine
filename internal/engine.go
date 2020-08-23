package internal

import (
	"github.com/shopspring/decimal"
	"sort"
	"strconv"
)

const (
	NotValidCoinAmount  = Error("Not valid amount coin inserted")
	NotEnoughCoinsErr   = Error("Not enough coins, please insert more")
	NotValidExchangeErr = Error("It is not possible to return exchange, please try with exact amount of coins")
)

type Error string

func (e Error) Error() string {
	return string(e)
}

type CashEngine struct {
	coinsForExchange      map[string]int
	currentServiceCoins   []Coin
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
		coinsForExchange:      storedCashDetailed,
		currentServiceCoins:   nil,
		currentServiceBalance: decimal.Decimal{},
		validCoins:            validCoins,
	}
}

// InsertCoins is used to load money in the machine for next buy. It allows user to insert any coin,
// but it could cause if it was not registered as a valid coin
func (e *CashEngine) InsertCoins(coins ...string) error {
	for _, c := range coins {
		newCoin, err := NewCoin(c)
		if err != nil {
			return err
		}
		if isNotValid(newCoin, e.validCoins) {
			return NotValidCoinAmount
		}
		e.currentServiceCoins = append(e.currentServiceCoins, newCoin)
		e.currentServiceBalance = e.currentServiceBalance.Add(newCoin.value)
	}

	return nil
}

func isNotValid(c Coin, validCoins []string) bool {
	for _, coin := range validCoins {
		coinFloat, err := strconv.ParseFloat(coin, 64)
		if err != nil {
			return false
		}
		if c.value.Equal(decimal.NewFromFloat(coinFloat)) {
			return false
		}
	}
	return true
}

func (e *CashEngine) DropCoins() []Coin {
	serviceCoins := e.currentServiceCoins
	e.currentServiceCoins = nil
	e.currentServiceBalance = decimal.Decimal{}

	return serviceCoins
}

// StoreCoins method is used to fill the machine with coins to let the vending machine have money to give exchange back
func (e *CashEngine) StoreCoins(coins ...Coin) {
	for _, c := range coins {
		e.coinsForExchange[c.category] += 1
	}
}

func (e *CashEngine) SellItem(price decimal.Decimal) ([]Coin, error) {
	if price.GreaterThan(e.currentServiceBalance) {
		return []Coin{}, NotEnoughCoinsErr
	}

	difference := e.currentServiceBalance.Sub(price)
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
			if e.coinsForExchange[currentCoin] >= neededCoins {
				for i := 0; i < neededCoins; i++ {
					exchange = append(exchange, coinReturned)
					e.coinsForExchange[currentCoin]--
				}

				e.currentServiceBalance = decimal.Zero
				e.currentServiceCoins = nil
				return exchange, nil
			}

			difference = difference.Sub(validCoinValue.Mul(decimal.NewFromInt(int64(e.coinsForExchange[currentCoin]))))
			for i := 0; i < e.coinsForExchange[currentCoin]; i++ {
				exchange = append(exchange, coinReturned)
				e.coinsForExchange[currentCoin]--
			}

		} else {
			var usedCoins int
			neededCoins := int(difference.Div(validCoinValue).IntPart())
			if e.coinsForExchange[currentCoin] >= neededCoins {
				usedCoins = neededCoins
			} else {
				usedCoins = e.coinsForExchange[currentCoin]
			}

			difference = difference.Sub(validCoinValue.Mul(decimal.NewFromInt(int64(usedCoins))))
			for i := 0; i < usedCoins; i++ {
				exchange = append(exchange, coinReturned)
				e.coinsForExchange[currentCoin]--
			}
		}
	}

	if difference.Equal(decimal.Zero) {
		e.currentServiceBalance = decimal.Zero
		e.currentServiceCoins = nil
		return exchange, nil
	}

	return []Coin{}, NotValidExchangeErr

}
