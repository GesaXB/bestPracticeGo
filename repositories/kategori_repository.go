package repositories

import (
	"ginBestPractice/models"

	"gorm.io/gorm"
)

type KategoriRepository interface {
	GetAll() ([]models.Kategori, error)
	GetById(id uint) (models.Kategori, error)
	Create(kategori models.Kategori) (models.Kategori, error)
}

type kategoriRepository struct {
	db *gorm.DB
}

func NewKategoriRepository(db *gorm.DB) KategoriRepository {
	return &kategoriRepository{db}
}

func (kr *kategoriRepository) GetAll() ([]models.Kategori, error) {
	var listKategori []models.Kategori
	err := kr.db.Find(&listKategori).Error
	return listKategori, err
}

func (kr *kategoriRepository) GetById(id uint) (models.Kategori, error) {
	var kategori models.Kategori
	err := kr.db.First(&kategori, id).Error
	return kategori, err
}

func (kr *kategoriRepository) Create(kategori models.Kategori) (models.Kategori, error) {
	err := kr.db.Create(&kategori).Error
	return kategori, err
}
