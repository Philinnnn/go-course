package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-course/internal/models"
	"go-course/internal/services"
	"net/http"
)

func RegisterTerminalRoutes(router *gin.RouterGroup) {
	router.POST("/", createTerminal)
	router.GET("/", getAllTerminals)
	router.GET("/:id", getTerminalByID)
	router.PUT("/:id", updateTerminal)
	router.DELETE("/:id", deleteTerminal)
}

// createTerminal godoc
// @Summary Создание нового терминала
// @Description Создаёт новый терминал с client_id, client_secret и uuid
// @Tags Терминалы
// @Accept json
// @Produce json
// @Param terminal body map[string]string true "Данные терминала"
// @Success 201 {object} models.Terminal
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /terminals/ [post]
func createTerminal(c *gin.Context) {
	var req struct {
		ClientID     string    `json:"client_id"`
		ClientSecret string    `json:"client_secret"`
		UUID         uuid.UUID `json:"uuid"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	terminal, err := services.CreateTerminal(req.ClientID, req.ClientSecret, req.UUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create terminal"})
		return
	}

	c.JSON(http.StatusCreated, terminal)
}

// getAllTerminals godoc
// @Summary Получение всех терминалов
// @Description Возвращает список всех терминалов
// @Tags Терминалы
// @Produce json
// @Success 200 {array} models.Terminal
// @Failure 500 {object} map[string]string
// @Router /terminals/ [get]
func getAllTerminals(c *gin.Context) {
	terminals, err := services.GetAllTerminals()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch terminals"})
		return
	}
	c.JSON(http.StatusOK, terminals)
}

// getTerminalByID godoc
// @Summary Получение терминала по ID
// @Description Возвращает терминал по его UUID
// @Tags Терминалы
// @Produce json
// @Param id path string true "UUID терминала"
// @Success 200 {object} models.Terminal
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /terminals/{id} [get]
func getTerminalByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	terminal, err := services.GetTerminalByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Terminal not found"})
		return
	}

	c.JSON(http.StatusOK, terminal)
}

// updateTerminal godoc
// @Summary Обновление данных терминала
// @Description Обновляет client_id, client_secret и uuid по ID
// @Tags Терминалы
// @Accept json
// @Param id path string true "UUID терминала"
// @Param terminal body map[string]string true "Обновлённые данные терминала"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /terminals/{id} [put]
func updateTerminal(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	var updated struct {
		ClientID     string    `json:"client_id"`
		ClientSecret string    `json:"client_secret"`
		UUID         uuid.UUID `json:"uuid"`
	}

	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err = services.UpdateTerminal(id, &models.Terminal{
		ClientID:     updated.ClientID,
		ClientSecret: updated.ClientSecret,
		UUID:         updated.UUID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update terminal"})
		return
	}

	c.Status(http.StatusNoContent)
}

// deleteTerminal godoc
// @Summary Удаление терминала
// @Description Удаляет терминал по ID
// @Tags Терминалы
// @Param id path string true "UUID терминала"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /terminals/{id} [delete]
func deleteTerminal(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	err = services.DeleteTerminal(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete terminal"})
		return
	}

	c.Status(http.StatusNoContent)
}
