package services

import (
	"socialVoleiAPI/internal/models"
	"socialVoleiAPI/internal/repositories"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id int) (*models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *userService) GetUserByID(id int) (*models.User, error) {
	return s.repo.GetByID(id)
}
