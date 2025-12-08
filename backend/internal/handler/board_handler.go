package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"trello-backend/internal/models"
	"trello-backend/internal/service"
)

type BoardHandler struct {
	boardService service.BoardService
}

func NewBoardHandler(boardService service.BoardService) *BoardHandler {
	return &BoardHandler{
		boardService: boardService,
	}
}

// GetBoards получает все доски пользователя
// GET /api/get-boards?user_id={user_id}
// Приоритет: user_id из JWT токена, если доступен, иначе из query параметра
func (h *BoardHandler) GetBoards(c *gin.Context) {
	var userID string

	// Пытаемся получить userID из JWT токена (если есть)
	if userIDFromContext, exists := c.Get("userID"); exists {
		var ok bool
		userID, ok = userIDFromContext.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID type"})
			return
		}
		// Используем userID из токена
	} else {
		// Если токена нет, получаем из query параметра
		userID = c.Query("user_id")
		if userID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
			return
		}
	}

	boards, err := h.boardService.GetBoardsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, boards)
}

// CreateBoard создает новую доску
// POST /api/save-board
func (h *BoardHandler) CreateBoard(c *gin.Context) {
	// Проверяем, что это не OPTIONS запрос (CORS preflight)
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	// Проверяем Content-Type только если он указан
	contentType := c.ContentType()
	if contentType != "" && !strings.HasPrefix(contentType, "application/json") {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Content-Type must be application/json",
			"received": contentType,
		})
		return
	}

	// Проверяем наличие тела запроса
	if c.Request.ContentLength == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "request body is required",
			"example": map[string]interface{}{
				"title":        "My Board",
				"backgroundId": "550e8400-e29b-41d4-a716-446655440000",
				"user_id":      "6949e844-9b4b-4dec-9d42-2d4216415f53",
			},
		})
		return
	}

	var req models.SaveBoardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// Более понятное сообщение об ошибке
		if err.Error() == "EOF" || strings.Contains(err.Error(), "EOF") {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "request body is empty or invalid JSON",
				"details": "Please send a valid JSON object with title, backgroundId, and user_id fields",
				"example": map[string]interface{}{
					"title":        "My Board",
					"backgroundId": "550e8400-e29b-41d4-a716-446655440000",
					"user_id":      "6949e844-9b4b-4dec-9d42-2d4216415f53",
				},
			})
			return
		}
		
		// Обработка ошибок валидации
		if validationErr, ok := err.(validator.ValidationErrors); ok {
			errors := make(map[string]string)
			for _, fieldErr := range validationErr {
				fieldName := fieldErr.Field()
				switch fieldName {
				case "BackgroundID":
					if strings.Contains(fieldErr.Tag(), "uuid") {
						errors["backgroundId"] = "backgroundId must be a valid UUID format (e.g., 550e8400-e29b-41d4-a716-446655440000)"
					} else if strings.Contains(fieldErr.Tag(), "required") {
						errors["backgroundId"] = "backgroundId is required"
					}
				case "UserID":
					if strings.Contains(fieldErr.Tag(), "uuid") {
						errors["user_id"] = "user_id must be a valid UUID format (e.g., 6949e844-9b4b-4dec-9d42-2d4216415f53)"
					} else if strings.Contains(fieldErr.Tag(), "required") {
						errors["user_id"] = "user_id is required"
					}
				case "Title":
					errors["title"] = "title is required and must be between 1 and 255 characters"
				default:
					errors[fieldName] = fieldErr.Error()
				}
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "validation failed",
				"details": errors,
				"example": map[string]interface{}{
					"title":        "My Board",
					"backgroundId": "550e8400-e29b-41d4-a716-446655440000",
					"user_id":      "6949e844-9b4b-4dec-9d42-2d4216415f53",
				},
			})
			return
		}
		
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid JSON format",
		})
		return
	}

	// Получаем userID из body запроса
	if req.UserID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required in request body"})
		return
	}

	if req.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
		return
	}

	if req.BackgroundID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "backgroundId is required"})
		return
	}

	board, err := h.boardService.CreateBoard(req.UserID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, board)
}

// ToggleFavorite переключает статус избранного для доски
// PATCH /api/save-favorite-board
func (h *BoardHandler) ToggleFavorite(c *gin.Context) {
	var req models.SaveFavoriteBoardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	boardID := req.BoardID
	if boardID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "board_id is required"})
		return
	}

	// Получаем userID из контекста
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID type"})
		return
	}

	board, err := h.boardService.ToggleFavorite(boardID, userIDStr)
	if err != nil {
		if err.Error() == "board not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if err.Error() == "access denied: board does not belong to user" {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, board)
}

