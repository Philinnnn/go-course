package migration

import (
	"go-course/internal/models"
	"gorm.io/gorm"
)

func SeedStatuses(db *gorm.DB) {
	statuses := []models.TransactionStatus{
		{"REFUND", "Возврат суммы"},
		{"AUTH", "Блокировка суммы"},
		{"CANCEL", "Разблокировка"},
		{"CHARGE", "Списание"},
		{"VERIFIED", "Проверка карты"},
		{"CANCEL_OLD", "Истёк срок CHARGE"},
		{"FAILED", "Неуспешно"},
		{"FINGERPRINT", "Проверка перед 3D"},
		{"3D", "Ошибка 3D"},
		{"NEW", "Ожидание"},
		{"REJECT", "Отклонено"},
	}

	for _, s := range statuses {
		db.FirstOrCreate(&models.TransactionStatus{}, s)
	}
}
