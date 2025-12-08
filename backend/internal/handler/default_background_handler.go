package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"trello-backend/internal/service"
)

type DefaultBackgroundHandler struct {
	backgroundService service.DefaultBackgroundService
}

func NewDefaultBackgroundHandler(backgroundService service.DefaultBackgroundService) *DefaultBackgroundHandler {
	return &DefaultBackgroundHandler{
		backgroundService: backgroundService,
	}
}

// GetDefaultBackgrounds получает все доступные фоновые изображения
// GET /api/default-backgrounds
func (h *DefaultBackgroundHandler) GetDefaultBackgrounds(c *gin.Context) {
	backgrounds, err := h.backgroundService.GetAllBackgrounds()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, backgrounds)
}

