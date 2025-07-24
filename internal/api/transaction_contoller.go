package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-course/internal/models"
	"go-course/internal/services"
	"net/http"
	"strconv"
	"time"
)

type createRequest struct {
	TerminalUUID string  `json:"terminal_uuid"`
	OrderID      string  `json:"order_id"`
	Amount       float64 `json:"amount"`
}

// CreateTransaction godoc
// @Summary Создать транзакцию
// @Description Создаёт новую транзакцию с указанием терминала, заказа и суммы
// @Tags Транзакции
// @Accept json
// @Produce json
// @Param transaction body createRequest true "Данные транзакции"
// @Success 201 {object} models.Transaction
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transactions [post]
func CreateTransaction(c *gin.Context) {
	var req createRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	terminalUUID, err := uuid.Parse(req.TerminalUUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid terminal_id"})
		return
	}

	tx, err := services.NewTransaction(terminalUUID, req.OrderID, req.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, tx)
}

// GetTransactionByID godoc
// @Summary Получить транзакцию по ID
// @Description Возвращает транзакцию по её ID
// @Tags Транзакции
// @Produce json
// @Param id path int true "ID транзакции"
// @Success 200 {object} models.Transaction
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /transactions/{id} [get]
func GetTransactionByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var tx models.Transaction
	if err := services.GetByID(id, &tx); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, tx)
}

// GetTransactionsByPeriod godoc
// @Summary Получить транзакции за период
// @Description Возвращает все транзакции между двумя датами (в формате RFC3339)
// @Tags Транзакции
// @Produce json
// @Param start query string true "Дата начала (RFC3339)"
// @Param end query string true "Дата окончания (RFC3339)"
// @Success 200 {array} models.Transaction
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transactions [get]
func GetTransactionsByPeriod(c *gin.Context) {
	startStr := c.Query("start")
	endStr := c.Query("end")

	start, err := time.Parse(time.RFC3339, startStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date"})
		return
	}
	end, err := time.Parse(time.RFC3339, endStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date"})
		return
	}

	txs, err := services.GetTransactionsByPeriod(start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get transactions"})
		return
	}

	c.JSON(http.StatusOK, txs)
}

type statusChangeRequest struct {
	Status string `json:"status"`
}

// ChangeTransactionStatus godoc
// @Summary Изменить статус транзакции
// @Description Меняет статус транзакции по правилам переходов
// @Tags Транзакции
// @Accept json
// @Produce json
// @Param id path int true "ID транзакции"
// @Param status body statusChangeRequest true "Новый статус транзакции"
// @Success 200 {object} models.Transaction
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /transactions/{id}/status [put]
func ChangeTransactionStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var req statusChangeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	var tx = &models.Transaction{}
	if err := services.GetByID(id, tx); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "transaction not found"})
		return
	}

	if err := services.ChangeStatus(tx, req.Status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tx)
}
