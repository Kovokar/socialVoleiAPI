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
	// mock temporário
	users := []models.User{
		{ID: 1, Name: "Alice", Email: "alice@mail.com"},
		{ID: 2, Name: "Bob", Email: "bob@mail.com"},
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
