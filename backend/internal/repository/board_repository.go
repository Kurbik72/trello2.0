package repository

import (
	"errors"
	"fmt"

	"trello-backend/internal/database"
	"trello-backend/internal/models"
)

type BoardRepository interface {
	Create(board *models.Board) error
	GetByID(id string) (*models.Board, error)
	GetByUserID(userID string) ([]models.Board, error)
	Update(board *models.Board) error
	Delete(id string) error
}

type boardRepository struct {
	db *database.DB
}

func NewBoardRepository(db *database.DB) BoardRepository {
	return &boardRepository{db: db}
}

func (r *boardRepository) Create(board *models.Board) error {
	if err := r.db.Create(board).Error; err != nil {
		return fmt.Errorf("failed to create board: %w", err)
	}
	return nil
}

func (r *boardRepository) GetByID(id string) (*models.Board, error) {
	var board models.Board
	if err := r.db.Where("id = ?", id).First(&board).Error; err != nil {
		return nil, fmt.Errorf("board not found: %w", err)
	}
	return &board, nil
}

func (r *boardRepository) GetByUserID(userID string) ([]models.Board, error) {
	var boards []models.Board
	if err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&boards).Error; err != nil {
		return nil, fmt.Errorf("failed to get boards: %w", err)
	}
	return boards, nil
}

func (r *boardRepository) Update(board *models.Board) error {
	result := r.db.Model(board).Updates(map[string]interface{}{
		"title":         board.Title,
		"is_favorite":   board.IsFavorite,
		"link_to_board": board.LinkToBoard,
		"background_src": board.BackgroundSrc,
		"background_id": board.BackgroundID,
	})
	
	if result.Error != nil {
		return fmt.Errorf("failed to update board: %w", result.Error)
	}
	
	if result.RowsAffected == 0 {
		return errors.New("board not found")
	}
	
	return nil
}

func (r *boardRepository) Delete(id string) error {
	result := r.db.Delete(&models.Board{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete board: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.New("board not found")
	}
	return nil
}

