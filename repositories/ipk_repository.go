package repositories

import (
	"ginBestPractice/models"

	"gorm.io/gorm"
)

type IpkRepositories interface {
	GetAll() ([]models.IPK, error)
}

type ipkRepositories struct {
	db *gorm.DB
}

func NewIpkRepository(db *gorm.DB) IpkRepositories {
	return &ipkRepositories{db}
}

func (r ipkRepositories) GetAll() ([]models.IPK, error) {
	var listIPk []models.IPK
	err := r.db.Find(&listIPk).Error
	return listIPk, err
}
