package services

import (
	"github.com/google/uuid"
	"go-course/internal/db"
	"go-course/internal/models"
	"time"
)

var allowedTransitions = map[string][]string{
	"NEW":    {"AUTH"},
	"AUTH":   {"CHARGE", "CANCEL"},
	"CHARGE": {"REFUND"},
}

func NewTransaction(terminalID uuid.UUID, orderID string, amount float64) (*models.Transaction, error) {
	tx := &models.Transaction{
		ID:            0,
		TerminalID:    terminalID,
		OrderID:       orderID,
		Amount:        amount,
		Status:        "NEW",
		CreatedAt:     time.Now(),
		StatusChanged: time.Now(),
	}

	if err := db.DB.Create(tx).Error; err != nil {
		return nil, err
	}
	return tx, nil
}

func ChangeStatus(transaction *models.Transaction, status string) error {
	if !isStatusChangeAvailable(transaction.Status, status) {
		return models.ErrStatusChangeNotAvailable
	}

	transaction.Status = status
	transaction.StatusChanged = time.Now()

	if err := db.DB.Save(transaction).Error; err != nil {
		return err
	}

	return nil
}

func isStatusChangeAvailable(from, to string) bool {
	for _, allowed := range allowedTransitions[from] {
		if allowed == to {
			return true
		}
	}
	return false
}

func GetByID(id uuid.UUID, tx *models.Transaction) error {
	return db.DB.First(tx, "id = ?", id).Error
}

func GetTransactionsByPeriod(start, end time.Time) ([]models.Transaction, error) {
	var txs []models.Transaction
	err := db.DB.Where("created_at BETWEEN ? AND ?", start, end).Find(&txs).Error
	return txs, err
}
