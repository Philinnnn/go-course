package models

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	TerminalID    uuid.UUID
	OrderID       string
	Amount        float64 `gorm:"type:numeric(12,2)"`
	Status        string  `gorm:"not null"`
	CreatedAt     time.Time
	StatusChanged time.Time
	Code          string
	Message       string

	Terminal  Terminal          `gorm:"foreignKey:TerminalID;constraint:OnDelete:CASCADE"`
	StatusRef TransactionStatus `gorm:"foreignKey:Status;references:Code;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}
