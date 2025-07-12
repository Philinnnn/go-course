package internal

import (
	"errors"

	"github.com/shopspring/decimal"
)

var ErrUnsupportedCurrency = errors.New("unsupported currency")

// Курсы валют на 12.07.2025
const (
	USD_KZT_rate = 522.37
	EUR_KZT_rate = 610.93
	RUB_KZT_rate = 6.70
)

func getKZTRate(currency string) (decimal.Decimal, error) {
	switch currency {
	case "USD":
		return decimal.NewFromFloat(USD_KZT_rate), nil
	case "EUR":
		return decimal.NewFromFloat(EUR_KZT_rate), nil
	case "RUB":
		return decimal.NewFromFloat(RUB_KZT_rate), nil
	default:
		return decimal.Zero, ErrUnsupportedCurrency
	}
}

func ConvertKZTTo(money decimal.Decimal, currency string) (decimal.Decimal, error) {
	rate, err := getKZTRate(currency)
	if err != nil {
		return decimal.Zero, err
	}
	return money.Div(rate).Round(2), nil
}

func ConvertToKZT(money decimal.Decimal, currency string) (decimal.Decimal, error) {
	rate, err := getKZTRate(currency)
	if err != nil {
		return decimal.Zero, err
	}
	return money.Mul(rate).Round(2), nil
}
