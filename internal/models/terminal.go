package models

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

var ErrUnauthorized = errors.New("unauthorized")

var jwtSecret = []byte("supersecretkey") // вынести в конфиг при необходимости

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

func (t *Terminal) CheckSecret(secret string) bool {
	return t.ClientSecret == secret
}

func (t *Terminal) GenerateJWT(secret string) (string, error) {
	if !t.CheckSecret(secret) {
		return "", ErrUnauthorized
	}
	return GenerateJWTForTerminal(t)
}

func GenerateJWTForTerminal(t *Terminal) (string, error) {
	claims := jwt.MapClaims{
		"terminal_id":   t.ID,
		"terminal_uuid": t.UUID.String(),
		"exp":           time.Now().Add(time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func JwtSecret() []byte {
	return jwtSecret
}
