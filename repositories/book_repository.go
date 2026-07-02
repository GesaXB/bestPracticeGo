package repositories

import (
	"ginBestPractice/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	GetAll() ([]models.Book, error)
	GetById(id uint) (models.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}

func (b bookRepository) GetAll() ([]models.Book, error) {
	var books []models.Book
	err := b.db.Find(&books).Error
	return books, err
}

func (b bookRepository) GetById(id uint) (models.Book, error) {
	var book models.Book
	err := b.db.First(&book, id).Error
	return book, err
}
