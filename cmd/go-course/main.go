package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"go-course/internal"
)

func main() {

	// Пример использования функции для конвертации валют
	// Конвертация 5000 KZT в USD
	converted, err := internal.ConvertKZT(decimal.NewFromFloat(5000), "to", "USD")
	if err == nil {
		fmt.Println("В USD:", converted)
	} else {
		fmt.Println("Ошибка:", err)
	}

	// Конвертация 17 EUR в KZT
	converted2, err2 := internal.ConvertKZT(decimal.NewFromFloat(17), "from", "EUR")
	if err2 == nil {
		fmt.Println("В KZT:", converted2)
	} else {
		fmt.Println("Ошибка:", err2)
	}

	// Пример использования функции для получения курса валюты
	transactions := []map[int]decimal.Decimal{
		{1: decimal.NewFromFloat(25000.00)},
		{1: decimal.NewFromFloat(20000.00)},
		{2: decimal.NewFromFloat(-9800.00)},
		{3: decimal.NewFromFloat(-1222.22)},
		{4: decimal.NewFromFloat(-1500.07)},
		{5: decimal.NewFromFloat(1201.37)},
		{6: decimal.NewFromFloat(-100.32)},
		{7: decimal.NewFromFloat(-523.33)},
	}

	internal.PrintWeeklySummary(transactions)
}
