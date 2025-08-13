package api

import (
	"github.com/gin-gonic/gin"
	"go-course/internal/services"
	"net/http"
)

// CurrencyConvert godoc
// @Summary Конвертация валют (демо)
// @Description Запускает конкурентный конвертер валют и возвращает результат
// @Tags Валюты
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /currency/convert [get]
func CurrencyConvert(c *gin.Context) {
	result := services.RunCurrencyConverterAPI()
	c.JSON(http.StatusOK, result)
}
