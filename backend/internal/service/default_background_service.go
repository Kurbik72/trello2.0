package service

import (
	"fmt"

	"trello-backend/internal/models"
	"trello-backend/internal/repository"
)

type DefaultBackgroundService interface {
	GetAllBackgrounds() ([]models.DefaultBackgroundResponse, error)
}

type defaultBackgroundService struct {
	backgroundRepo repository.DefaultBackgroundRepository
}

func NewDefaultBackgroundService(backgroundRepo repository.DefaultBackgroundRepository) DefaultBackgroundService {
	return &defaultBackgroundService{
		backgroundRepo: backgroundRepo,
	}
}

func (s *defaultBackgroundService) GetAllBackgrounds() ([]models.DefaultBackgroundResponse, error) {
	backgrounds, err := s.backgroundRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get backgrounds: %w", err)
	}

	return models.DefaultBackgroundsToResponse(backgrounds), nil
}

