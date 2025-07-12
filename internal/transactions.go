package internal

import (
	"fmt"

	"github.com/shopspring/decimal"
)

var weekDays = [...]string{
	"Неизвестный день",
	"Понедельник",
	"Вторник",
	"Среда",
	"Четверг",
	"Пятница",
	"Суббота",
	"Воскресенье",
}

func PrintWeeklySummary(transactions []map[int]decimal.Decimal) {
	total := decimal.Zero

	for _, tx := range transactions {
		for day, amount := range tx {
			dayName := weekDays[0]
			if day >= 1 && day <= 7 {
				dayName = weekDays[day]
			}

			fmt.Println(dayName)
			if amount.GreaterThan(decimal.Zero) {
				fmt.Printf("Поступление: %s\n\n", amount.StringFixed(2))
			} else {
				fmt.Printf("Списание: %s\n\n", amount.Abs().StringFixed(2))
			}

			total = total.Add(amount)
		}
	}

	fmt.Printf("Итог за неделю: %s\n", total.StringFixed(2))
}
