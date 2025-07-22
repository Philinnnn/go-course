package models

import "github.com/google/uuid"

type Terminal struct {
	ID           uint64 `gorm:"primaryKey,autoIncrement"`
	ClientID     string
	ClientSecret string
	UUID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
}
