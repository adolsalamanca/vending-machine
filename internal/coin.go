package internal

import "github.com/shopspring/decimal"

const (
	FiveCent       = "0.05"
	TenCent        = "0.10"
	TwentyFiveCent = "0.25"
	OneUnit        = "1.00"
	TwoUnit        = "2.00"
)

type Coin struct {
	category string
	value    decimal.Decimal
}

func NewCoin(t string) (Coin, error) {
	d, err := decimal.NewFromString(t)
	if err != nil {
		return Coin{}, err
	}
	return Coin{
		category: t,
		value:    d,
	}, nil
}

func (c *Coin) Value() decimal.Decimal {
	return c.value
}
