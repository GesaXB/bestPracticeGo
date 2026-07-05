package repositories

import (
	"ginBestPractice/models"

	"gorm.io/gorm"
)

type MakananRepository interface {
	GetAll() ([]models.Makanan, error)
	GetById(id uint) (models.Makanan, error)
	Create(makanan models.Makanan) (models.Makanan, error)
	Update(makanan models.Makanan) (models.Makanan, error)
}

type makananRepository struct {
	db *gorm.DB
}

func NewMakananRepository(db *gorm.DB) MakananRepository {
	return &makananRepository{db}
}

func (r *makananRepository) GetAll() ([]models.Makanan, error) {
	var listMakanan []models.Makanan
	err := r.db.Find(&listMakanan).Error
	return listMakanan, err
}

func (r *makananRepository) GetById(id uint) (models.Makanan, error) {
	var makanan models.Makanan
	err := r.db.First(&makanan, id).Error
	return makanan, err
}

func (r *makananRepository) Create(makanan models.Makanan) (models.Makanan, error) {
	err := r.db.Create(&makanan).Error
	return makanan, err
}

func (r *makananRepository) Update(makanan models.Makanan) (models.Makanan, error) {
	err := r.db.Save(&makanan).Error
	return makanan, err
}
