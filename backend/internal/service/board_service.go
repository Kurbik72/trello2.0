package service

import (
	"fmt"
	"strings"

	"trello-backend/internal/models"
	"trello-backend/internal/repository"
)

type BoardService interface {
	CreateBoard(userID string, req *models.SaveBoardRequest) (*models.BoardResponse, error)
	GetBoardsByUserID(userID string) ([]models.BoardResponse, error)
	ToggleFavorite(boardID string, userID string) (*models.BoardResponse, error)
	GetBoardByID(boardID string, userID string) (*models.BoardResponse, error)
	DeleteBoard(boardID string, userID string) error
}

type boardService struct {
	boardRepo           repository.BoardRepository
	backgroundRepo      repository.DefaultBackgroundRepository
}

func NewBoardService(
	boardRepo repository.BoardRepository,
	backgroundRepo repository.DefaultBackgroundRepository,
) BoardService {
	return &boardService{
		boardRepo:      boardRepo,
		backgroundRepo: backgroundRepo,
	}
}

func (s *boardService) CreateBoard(userID string, req *models.SaveBoardRequest) (*models.BoardResponse, error) {
	backgroundID := req.BackgroundID

	// Получаем фоновое изображение
	background, err := s.backgroundRepo.GetByID(backgroundID)
	if err != nil {
		return nil, fmt.Errorf("background not found: %w", err)
	}

	// Генерируем ссылку на доску
	linkToBoard := s.generateBoardLink()

	board := &models.Board{
		Title:         strings.TrimSpace(req.Title),
		IsFavorite:    false,
		LinkToBoard:   linkToBoard,
		BackgroundSrc: background.Src,
		UserID:        userID,
		BackgroundID:  &backgroundID,
	}

	if err := s.boardRepo.Create(board); err != nil {
		return nil, fmt.Errorf("failed to create board: %w", err)
	}

	return models.BoardToResponse(board), nil
}

func (s *boardService) GetBoardsByUserID(userID string) ([]models.BoardResponse, error) {
	boards, err := s.boardRepo.GetByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get boards: %w", err)
	}

	return models.BoardsToResponse(boards), nil
}

func (s *boardService) ToggleFavorite(boardID string, userID string) (*models.BoardResponse, error) {
	board, err := s.boardRepo.GetByID(boardID)
	if err != nil {
		return nil, fmt.Errorf("board not found: %w", err)
	}

	// Проверяем, что доска принадлежит пользователю
	if board.UserID != userID {
		return nil, fmt.Errorf("access denied: board does not belong to user")
	}

	// Переключаем статус избранного
	board.IsFavorite = !board.IsFavorite

	if err := s.boardRepo.Update(board); err != nil {
		return nil, fmt.Errorf("failed to update board: %w", err)
	}

	return models.BoardToResponse(board), nil
}

func (s *boardService) GetBoardByID(boardID string, userID string) (*models.BoardResponse, error) {
	board, err := s.boardRepo.GetByID(boardID)
	if err != nil {
		return nil, fmt.Errorf("board not found: %w", err)
	}

	// Проверяем, что доска принадлежит пользователю
	if board.UserID != userID {
		return nil, fmt.Errorf("access denied: board does not belong to user")
	}

	return models.BoardToResponse(board), nil
}

func (s *boardService) DeleteBoard(boardID string, userID string) error {
	board, err := s.boardRepo.GetByID(boardID)
	if err != nil {
		return fmt.Errorf("board not found: %w", err)
	}

	// Проверяем, что доска принадлежит пользователю
	if board.UserID != userID {
		return fmt.Errorf("access denied: board does not belong to user")
	}

	if err := s.boardRepo.Delete(boardID); err != nil {
		return fmt.Errorf("failed to delete board: %w", err)
	}

	return nil
}

func (s *boardService) generateBoardLink() string {
	// Генерируем уникальную ссылку на доску
	// UUID будет сгенерирован в Prisma при создании записи
	return "/board/temp"
}

