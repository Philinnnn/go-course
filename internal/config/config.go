package config

import (
	"github.com/spf13/viper"
	"go-course/pkg/logger"
	"strconv"
)

type AppConfig struct {
	IP          string
	Port        int
	AutoMigrate bool
	DB          DBConfig
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

var Config AppConfig

func Load() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		logger.Warn(".env file not found")
	}

	Config = AppConfig{
		IP:   viper.GetString("APP_IP"),
		Port: viper.GetInt("APP_PORT"),
		DB: DBConfig{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetInt("DB_PORT"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASS"),
			Name:     viper.GetString("DB_NAME"),
		},
	}

	Config.AutoMigrate, _ = strconv.ParseBool(viper.GetString("AUTO_MIGRATE"))
}
