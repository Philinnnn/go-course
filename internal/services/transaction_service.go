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

func NewTransaction(terminalUUID uuid.UUID, orderID string, amount float64) (*models.Transaction, error) {
	tx := &models.Transaction{
		TerminalUUID:  terminalUUID,
		OrderID:       orderID,
		Amount:        amount,
		Status:        "NEW",
		CreatedAt:     time.Now(),
		StatusChanged: time.Now(),
	}

	if err := db.DB.Create(tx).Error; err != nil {
		return nil, err
	}

	if err := db.DB.Preload("Terminal").First(tx, tx.ID).Error; err != nil {
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

func GetByID(id uint64, tx *models.Transaction) error {
	err := db.DB.Preload("Terminal").First(tx, "id = ?", id).Error
	if err != nil {
		return err
	}
	if tx.TerminalUUID != uuid.Nil && tx.Terminal.ID == 0 {
		var term models.Terminal
		db.DB.First(&term, "uuid = ?", tx.TerminalUUID)
		tx.Terminal = term
	}
	return nil
}

func GetTransactionsByPeriod(start, end time.Time) ([]models.Transaction, error) {
	var txs []models.Transaction
	err := db.DB.Preload("Terminal").Where("created_at BETWEEN ? AND ?", start, end).Find(&txs).Error
	return txs, err
}
