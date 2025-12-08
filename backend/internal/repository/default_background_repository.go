package repository

import (
	"fmt"

	"trello-backend/internal/database"
	"trello-backend/internal/models"
)

type DefaultBackgroundRepository interface {
	GetAll() ([]models.DefaultBackground, error)
	GetByID(id string) (*models.DefaultBackground, error)
	Create(background *models.DefaultBackground) error
}

type defaultBackgroundRepository struct {
	db *database.DB
}

func NewDefaultBackgroundRepository(db *database.DB) DefaultBackgroundRepository {
	return &defaultBackgroundRepository{db: db}
}

func (r *defaultBackgroundRepository) GetAll() ([]models.DefaultBackground, error) {
	var backgrounds []models.DefaultBackground
	if err := r.db.Order("created_at ASC").Find(&backgrounds).Error; err != nil {
		return nil, fmt.Errorf("failed to get default backgrounds: %w", err)
	}
	return backgrounds, nil
}

func (r *defaultBackgroundRepository) GetByID(id string) (*models.DefaultBackground, error) {
	var background models.DefaultBackground
	if err := r.db.Where("id = ?", id).First(&background).Error; err != nil {
		return nil, fmt.Errorf("default background not found: %w", err)
	}
	return &background, nil
}

func (r *defaultBackgroundRepository) Create(background *models.DefaultBackground) error {
	if err := r.db.Create(background).Error; err != nil {
		return fmt.Errorf("failed to create default background: %w", err)
	}
	return nil
}

