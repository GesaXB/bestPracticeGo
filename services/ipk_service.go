package services

import (
	"ginBestPractice/models"
	"ginBestPractice/repositories"
)

type IpkService interface {
	GetAll() ([]models.IPK, error)
}

type ipkService struct {
	repo repositories.IpkRepositories
}

func NewIpkService(repo repositories.IpkRepositories) IpkService {
	return &ipkService{repo}
}

func (i ipkService) GetAll() ([]models.IPK, error) {
	return i.repo.GetAll()
}
