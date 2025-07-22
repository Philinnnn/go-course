package models

import (
	"errors"
)

var ErrStatusChangeNotAvailable = errors.New("status change not available")

type TransactionStatus struct {
	Status      string `gorm:"primaryKey"`
	Description string
}
