package db

import (
	"fmt"
	"go-course/internal/config"
	"go-course/internal/migration"
	"go-course/internal/models"
	"go-course/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dbConfig := config.Config.DB

	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Name,
	)

	database, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		logger.Fatal("Database connection error: ", err)
	}

	sqlDB, err := database.DB()
	if err != nil {
		logger.Fatal("SQL DB error: ", err)
	}

	if err := sqlDB.Ping(); err != nil {
		logger.Fatal("DB unreachable: ", err)
	}

	DB = database
	logger.Success("DB connection established successfully")

	if config.Config.AutoMigrate {
		err := DB.AutoMigrate(
			&models.Terminal{},
			&models.Transaction{},
			&models.TransactionStatus{},
		)
		if err != nil {
			logger.Error("AutoMigrate error")
			return
		}
		migration.SeedStatuses(DB)
		logger.Success("Migrated successfully")
	}
}
