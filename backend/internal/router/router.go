package router

import (
	"trello-backend/internal/config"
	"trello-backend/internal/handler"
	"trello-backend/internal/middleware"
	"trello-backend/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	cfg *config.Config,
	boardHandler *handler.BoardHandler,
	backgroundHandler *handler.DefaultBackgroundHandler,
) *gin.Engine {
	gin.SetMode(cfg.Server.GinMode)

	r := gin.Default()

	// Middleware
	r.Use(middleware.CORSMiddleware())
	r.Use(gin.Recovery())

	// JWT Manager
	jwtManager := jwt.NewJWTManager(cfg.JWT.Secret, cfg.JWT.GetAccessTokenExpiration())

	// API routes
	api := r.Group("/api")
	{
		// Public routes
		api.GET("/default-backgrounds", backgroundHandler.GetDefaultBackgrounds)
		// Временно убрана авторизация для save-board
		api.POST("/save-board", boardHandler.CreateBoard)

		// Protected routes (require authentication)
		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware(jwtManager))
		{
			protected.PATCH("/save-favorite-board", boardHandler.ToggleFavorite)
		}

		// Routes with optional authentication (for backward compatibility)
		api.GET("/get-boards", middleware.OptionalAuthMiddleware(jwtManager), boardHandler.GetBoards)
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return r
}

