package services

import (
	"fmt"
	"socialVoleiAPI/internal/models"
	"socialVoleiAPI/internal/repositories"
	"strconv"
)

type EstablishmentService struct {
	repo *repositories.EstablishmentRepository
}

func NewEstablishmentService(repo *repositories.EstablishmentRepository) *EstablishmentService {
	return &EstablishmentService{repo: repo}
}

func (s *EstablishmentService) CreateEstablishment(req *models.Establishment) error {

	if req.Name == "" {
		return fmt.Errorf("Nome é obrigatório")
	}
	if req.CNPJ == "" {
		return fmt.Errorf("CNPJ é obrigatório")
	}

	Establishment := &models.Establishment{
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		CNPJ:      req.CNPJ,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}

	return s.repo.CreateEstablishment(Establishment)
}

func (s *EstablishmentService) GetAllEstablishments() ([]models.Establishment, error) {
	return s.repo.FindAllEstablishments()
}

func (s *EstablishmentService) GetEstablishmentByID(id string) (models.Establishment, error) {

	var estab models.Establishment
	intId, err := strconv.Atoi(id)
	if err != nil {
		return estab, fmt.Errorf("erro ao converter id para inteiro")
	}
	return s.repo.FindEstablishmentByID(estab, intId)
}
