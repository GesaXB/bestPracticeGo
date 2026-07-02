package repositories

import (
	"ginBestPractice/models"

	"gorm.io/gorm"
)

type MahasiswaRepository interface {
	FindAll() ([]models.Mahasiswa, error)
	FindById(id uint) (models.Mahasiswa, error)
}

type mahasiswaRepository struct {
	db *gorm.DB
}

func NewMahasiswaRepository(db *gorm.DB) MahasiswaRepository {
	return &mahasiswaRepository{db}
}

func (r *mahasiswaRepository) FindAll() ([]models.Mahasiswa, error) {
	var mhs []models.Mahasiswa
	err := r.db.Find(&mhs).Error
	return mhs, err
}

func (r *mahasiswaRepository) FindById(id uint) (models.Mahasiswa, error) {
	var mhs models.Mahasiswa
	err := r.db.First(&mhs, id).Error
	return mhs, err
}
