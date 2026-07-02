package services

import (
	"ginBestPractice/models"
	"ginBestPractice/repositories"
)

type BookService interface {
	GetAll() ([]models.Book, error)
	GetByID(id uint) (models.Book, error)
}

type bookService struct {
	repo repositories.BookRepository
}

func NewBookService(repo repositories.BookRepository) BookService {
	return &bookService{repo}
}

func (bs bookService) GetAll() ([]models.Book, error) {
	return bs.repo.GetAll()
}

func (bs bookService) GetByID(id uint) (models.Book, error) {
	return bs.repo.GetById(id)
}
