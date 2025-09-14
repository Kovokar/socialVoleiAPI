package repositories

import "socialVoleiAPI/internal/models"

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetByID(id int) (*models.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) GetAll() ([]models.User, error) {
	users := []models.User{
		{
			ID:            1,
			Name:          "Alice",
			Email:         "alice@mail.com",
			Telefone:      "(11) 99999-0001",
			Longitude:     -46.6333,
			Latitude:      -23.5505,
			Senha:         "senha123",
			CodigoConvite: "ABC123",
			Sexo:          "F",
			Categoria:     "comum",
		},
		{
			ID:            2,
			Name:          "Bob",
			Email:         "bob@mail.com",
			Telefone:      "(11) 99999-0002",
			Longitude:     -46.6388,
			Latitude:      -23.5587,
			Senha:         "senha456",
			CodigoConvite: "XYZ789",
			Sexo:          "M",
			Categoria:     "premium",
		},
	}
	return users, nil
}

func (r *userRepository) GetByID(id int) (*models.User, error) {
	users, _ := r.GetAll()
	for _, u := range users {
		if u.ID == id {
			return &u, nil
		}
	}
	return nil, nil
}
