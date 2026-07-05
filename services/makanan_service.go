package services

import (
	"errors"
	"ginBestPractice/dtos"
	"ginBestPractice/models"
	"ginBestPractice/repositories"

	"gorm.io/gorm"
)

type MakananService interface {
	GetAll() ([]models.Makanan, error)
	GetById(id uint) (models.Makanan, error)
	Create(input dtos.MakananCreateRequest) (dtos.MakananResponse, error)
	Update(id uint, input dtos.MakananUpdateRequest) (dtos.MakananResponse, error)
}

type makananService struct {
	repo repositories.MakananRepository
}

func NewMakananService(repo repositories.MakananRepository) MakananService {
	return &makananService{repo}
}

func (s *makananService) GetAll() ([]models.Makanan, error) {
	return s.repo.GetAll()
}

func (s *makananService) GetById(id uint) (models.Makanan, error) {
	return s.repo.GetById(id)
}

func (s *makananService) Create(input dtos.MakananCreateRequest) (dtos.MakananResponse, error) {
	makanan := models.Makanan{
		Name:        input.Name,
		Description: input.Name,
	}

	res, err := s.repo.Create(makanan)
	if err != nil {
		return dtos.MakananResponse{}, err
	}

	return dtos.MakananResponse{
		Id:          res.Id,
		Name:        res.Name,
		Description: res.Description,
	}, nil
}

func (s *makananService) Update(id uint, input dtos.MakananUpdateRequest) (dtos.MakananResponse, error) {
	makanan, err := s.repo.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dtos.MakananResponse{}, errors.New("Data tidak bisa ditemukan")
		}
		return dtos.MakananResponse{}, err
	}

	makanan.Name = input.Name
	makanan.Description = input.Description

	updateMakanan, err := s.repo.Update(makanan)
	if err != nil {
		return dtos.MakananResponse{}, err
	}

	return dtos.MakananResponse{
		Id:          updateMakanan.Id,
		Name:        updateMakanan.Name,
		Description: updateMakanan.Description,
	}, nil
}
