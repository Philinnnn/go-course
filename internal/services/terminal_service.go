package services

import (
	"github.com/google/uuid"
	"go-course/internal/db"
	"go-course/internal/models"
)

func CreateTerminal(clientID, clientSecret string, uuid uuid.UUID) (*models.Terminal, error) {
	t := &models.Terminal{
		ID:           0,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		UUID:         uuid,
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

func GetTerminalByID(id uuid.UUID) (*models.Terminal, error) {
	var terminal models.Terminal
	if err := db.DB.First(&terminal, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &terminal, nil
}

func UpdateTerminal(id uuid.UUID, updated *models.Terminal) error {
	return db.DB.Model(&models.Terminal{}).Where("id = ?", id).Updates(updated).Error
}

func DeleteTerminal(id uuid.UUID) error {
	return db.DB.Delete(&models.Terminal{}, "id = ?", id).Error
}
