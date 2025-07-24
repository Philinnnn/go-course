package services

import (
	"github.com/google/uuid"
	"go-course/internal/db"
	"go-course/internal/models"
)

func CreateTerminal(clientID, clientSecret string) (*models.Terminal, error) {
	t := &models.Terminal{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
	if err := db.DB.Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func GetAllTerminals() ([]models.Terminal, error) {
	var terminals []models.Terminal
	if err := db.DB.Find(&terminals).Error; err != nil {
		return nil, err
	}
	return terminals, nil
}

func GetTerminalByUUID(uuidVal uuid.UUID) (*models.Terminal, error) {
	var terminal models.Terminal
	if err := db.DB.First(&terminal, "uuid = ?", uuidVal).Error; err != nil {
		return nil, err
	}
	return &terminal, nil
}

func UpdateTerminal(uuidVal uuid.UUID, updated *models.Terminal) error {
	return db.DB.Model(&models.Terminal{}).Where("uuid = ?", uuidVal).Updates(updated).Error
}

func DeleteTerminal(uuidVal uuid.UUID) error {
	return db.DB.Delete(&models.Terminal{}, "uuid = ?", uuidVal).Error
}
