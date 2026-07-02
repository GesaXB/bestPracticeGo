package services

import (
	"ginBestPractice/models"
	"ginBestPractice/repositories"
)

type MakananService interface {
	GetById(id uint) (models.Makanan, error)
}

type makananService struct {
	repo repositories.MakananRepository
}

func NewMakananService(repo repositories.MakananRepository) MakananService {
	return &makananService{repo}
}

func (s *makananService) GetById(id uint) (models.Makanan, error) {
	return s.repo.GetById(id)
}
