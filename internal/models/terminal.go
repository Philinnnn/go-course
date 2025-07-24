package models

import (
	"github.com/google/uuid"
)

type Terminal struct {
	ID           uint `gorm:"primaryKey;autoIncrement"`
	ClientID     string
	ClientSecret string
	UUID         uuid.UUID `gorm:"type:uuid;uniqueIndex"`
}

func (t *Terminal) BeforeCreate() (err error) {
	if t.UUID == uuid.Nil {
		t.UUID = uuid.New()
	}
	return
}
