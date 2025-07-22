package models

import "github.com/google/uuid"

type Terminal struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	ClientID     string
	ClientSecret string
	UUID         string
}
