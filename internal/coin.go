package internal

import "github.com/shopspring/decimal"

const (
	NotValidCoin   = Error("Not valid coin inserted")
	FiveCent       = "0.05"
	TenCent        = "0.1"
	TwentyFiveCent = "0.25"
	OneUnit        = "1"
	TwoUnit        = "2"
)

type Coin struct {
	category string
	value    decimal.Decimal
}

// NewCoin instantiates a coin given a string
func NewCoin(t string) (Coin, error) {
	d, err := decimal.NewFromString(t)
	if err != nil {
		return Coin{}, NotValidCoin
	}
	return Coin{
		category: t,
		value:    d,
	}, nil
}

func (c *Coin) Value() decimal.Decimal {
	return c.value
}
