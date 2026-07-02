package repositories

import (
	"ginBestPractice/models"

	"gorm.io/gorm"
)

type MakananRepository interface {
	GetById(id uint) (models.Makanan, error)
}

type makananRepository struct {
	db *gorm.DB
}

func NewMakananRepository(db *gorm.DB) MakananRepository {
	return &makananRepository{db}
}

func (r *makananRepository) GetById(id uint) (models.Makanan, error) {
	var makanan models.Makanan
	err := r.db.First(&makanan, id).Error
	return makanan, err
}
