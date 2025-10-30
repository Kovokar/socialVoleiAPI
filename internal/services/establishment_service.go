package services

import (
	"fmt"
	"socialVoleiAPI/internal/models"
	"socialVoleiAPI/internal/repositories"
	"socialVoleiAPI/internal/utils"
	"socialVoleiAPI/internal/utils/validations"
	"strconv"
)

type EstablishmentService struct {
	repo *repositories.EstablishmentRepository
}

func NewEstablishmentService(repo *repositories.EstablishmentRepository) *EstablishmentService {
	return &EstablishmentService{repo: repo}
}

func (s *EstablishmentService) CreateEstablishment(req *models.Establishment) error {

	if err := validations.ValidateRequiredFields(
		validations.Field{Name: "Nome", Value: req.Name},
		validations.Field{Name: "CNPJ", Value: req.CNPJ},
		validations.Field{Name: "Email", Value: req.Email},
	); err != nil {
		return err
	}

	formatCNPJ, formatPhone := utils.RmMaskCNPJ(req.CNPJ), utils.RmMaskPhone(req.Phone)
	fmt.Println(formatCNPJ, formatPhone)

	Establishment := &models.Establishment{
		Name:      req.Name,
		Email:     req.Email,
		Phone:     formatPhone,
		CNPJ:      formatCNPJ,
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
