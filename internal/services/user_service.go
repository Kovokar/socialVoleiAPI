package services

import (
	"socialVoleiAPI/internal/models"
	"socialVoleiAPI/internal/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(req *models.RegisterRequest) (*models.User, error) {
	//
	// // // aqui implementar a logica de hashear a senha
	//

	user := &models.User{
		Name:  req.Name,
		Email: req.Email,
		// Password: // hash,
		Phone:     req.Phone,
		Gender:    models.GenderType(req.Gender),
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		// InviterId: req.Invite, // logica de pegar o id
	}

	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}
	return user, nil
}
