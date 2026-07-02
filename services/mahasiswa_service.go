package services

import (
	"ginBestPractice/models"
	"ginBestPractice/repositories"
)

type MahasiswaService interface {
	GetAll() ([]models.Mahasiswa, error)
	GetByID(id uint) (models.Mahasiswa, error)
}

type mahasiswaService struct {
	repo repositories.MahasiswaRepository
}

func NewMahasiswaService(repo repositories.MahasiswaRepository) MahasiswaService {
	return &mahasiswaService{repo}
}

func (s *mahasiswaService) GetAll() ([]models.Mahasiswa, error) {
	return s.repo.FindAll()
}

func (s *mahasiswaService) GetByID(id uint) (models.Mahasiswa, error) {
	return s.repo.FindById(id)
}
