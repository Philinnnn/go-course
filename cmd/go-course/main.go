package main

import (
	"fmt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go-course/docs"
	"go-course/internal/config"
	"go-course/internal/db"
	"go-course/internal/routes"
)

// @title Transaction API
// @version 1.0
// @description Сервис учёта транзакций мерчантов
// @host localhost:8080
// @BasePath /api
func main() {
	config.Load()
	db.Init()

	router := routes.SetupRouter()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		panic(fmt.Errorf("failed to set trusted proxies: %w", err))
	}

	addr := fmt.Sprintf("%s:%d", config.Config.IP, config.Config.Port)
	if err := router.Run(addr); err != nil {
		panic(fmt.Errorf("failed to run server: %w", err))
	}
}
