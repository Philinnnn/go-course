package routes

import (
	"github.com/gin-gonic/gin"
	"go-course/internal/api"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/transactions/:id", api.GetTransactionByID)
		apiGroup.GET("/transactions", api.GetTransactionsByPeriod)
		apiGroup.POST("/transactions", api.CreateTransaction)
		apiGroup.PUT("/transactions/:id/status", api.ChangeTransactionStatus)

		terminalsGroup := apiGroup.Group("/terminals")
		api.RegisterTerminalRoutes(terminalsGroup)
	}

	return r
}
