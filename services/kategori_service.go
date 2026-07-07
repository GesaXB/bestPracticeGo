package services

import (
	"ginBestPractice/dtos"
	"ginBestPractice/models"
	"ginBestPractice/repositories"
)

type KategoriService interface {
	GetAll() ([]models.Kategori, error)
	GetById(id uint) (models.Kategori, error)
	Create(input dtos.KategoriCreateRequest) (dtos.KategoriResponse, error)
}

type kategoriService struct {
	repo repositories.KategoriRepository
}

func NewKategoriService(repo repositories.KategoriRepository) KategoriService {
	return &kategoriService{repo}
}

func (ks *kategoriService) GetAll() ([]models.Kategori, error) {
	return ks.repo.GetAll()
}

func (ks *kategoriService) GetById(id uint) (models.Kategori, error) {
	return ks.repo.GetById(id)
}

func (ks kategoriService) Create(input dtos.KategoriCreateRequest) (dtos.KategoriResponse, error) {
	kategori := models.Kategori{
		Name: input.Name,
	}

	res, err := ks.repo.Create(kategori)
	if err != nil {
		return dtos.KategoriResponse{}, err
	}

	return dtos.KategoriResponse{
		Name: res.Name,
	}, nil
}
