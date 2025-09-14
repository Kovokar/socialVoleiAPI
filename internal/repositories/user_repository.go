package repositories

import "socialVoleiAPI/internal/models"

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetByID(id int) (*models.User, error)
	Create(user models.User) (*models.User, error)
}

type userRepository struct {
	users  []models.User
	nextID int
}

func NewUserRepository() UserRepository {
	// ✅ Carrega os dados mockados apenas uma vez, na criação do repositório
	mockedUsers := []models.User{
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

	return &userRepository{
		users:  mockedUsers,
		nextID: 3, // Já temos 2 usuários mockados
	}
}

func (r *userRepository) GetAll() ([]models.User, error) {
	return r.users, nil
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

func (r *userRepository) Create(user models.User) (*models.User, error) {
	user.ID = r.nextID
	r.nextID++
	r.users = append(r.users, user)
	return &user, nil
}
