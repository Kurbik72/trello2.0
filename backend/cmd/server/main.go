package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"trello-backend/internal/config"
	"trello-backend/internal/database"
	"trello-backend/internal/handler"
	"trello-backend/internal/repository"
	"trello-backend/internal/router"
	"trello-backend/internal/service"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Подключение к базе данных
	db, err := database.New(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Инициализация репозиториев
	boardRepo := repository.NewBoardRepository(db)
	backgroundRepo := repository.NewDefaultBackgroundRepository(db)
	// userRepo := repository.NewUserRepository(db) // Для будущей реализации аутентификации

	// Инициализация сервисов
	boardService := service.NewBoardService(boardRepo, backgroundRepo)
	backgroundService := service.NewDefaultBackgroundService(backgroundRepo)

	// Инициализация хендлеров
	boardHandler := handler.NewBoardHandler(boardService)
	backgroundHandler := handler.NewDefaultBackgroundHandler(backgroundService)

	// Настройка роутера
	r := router.SetupRouter(cfg, boardHandler, backgroundHandler)

	// Создание HTTP сервера
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Server.Port),
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Запуск сервера в горутине
	go func() {
		log.Printf("Server starting on port %s", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}

