package models

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement"`
	TerminalUUID  uuid.UUID `gorm:"type:uuid;not null"`
	OrderID       string
	Amount        float64 `gorm:"type:numeric(12,2)"`
	Status        string  `gorm:"not null"`
	CreatedAt     time.Time
	StatusChanged time.Time
	Code          string
	Message       string

	Terminal Terminal `gorm:"foreignKey:TerminalUUID;references:UUID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}
