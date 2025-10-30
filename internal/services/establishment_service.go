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

func (s *EstablishmentService) BulkCreateEstablishment(req []models.Establishment) error {
	var establishments []models.Establishment

	for _, val := range req {
		// Validação de campos obrigatórios
		if err := validations.ValidateRequiredFields(
			validations.Field{Name: "Nome", Value: val.Name},
			validations.Field{Name: "CNPJ", Value: val.CNPJ},
			validations.Field{Name: "Email", Value: val.Email},
		); err != nil {
			return fmt.Errorf("erro no estabelecimento %s: %w", val.Name, err)
		}

		formatCNPJ := utils.RmMaskCNPJ(val.CNPJ)
		formatPhone := utils.RmMaskPhone(val.Phone)

		e := models.Establishment{
			Name:      val.Name,
			Email:     val.Email,
			Phone:     formatPhone,
			CNPJ:      formatCNPJ,
			Latitude:  val.Latitude,
			Longitude: val.Longitude,
		}

		establishments = append(establishments, e)
	}

	// Envia o slice para o repositório
	return s.repo.BulkCreateEstablishment(&establishments)
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

func (s *EstablishmentService) UpdateEstablishment(id string, req models.Establishment) error {
	// Valida campos obrigatórios

	formatCNPJ := utils.RmMaskCNPJ(req.CNPJ)
	formatPhone := utils.RmMaskPhone(req.Phone)

	updateData := models.Establishment{
		Name:      req.Name,
		Email:     req.Email,
		Phone:     formatPhone,
		CNPJ:      formatCNPJ,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}

	return s.repo.UpdateEstablishment(id, updateData)
}

func (s *EstablishmentService) DeleteEstablishment(id string) error {
	if id == "" {
		return fmt.Errorf("ID é obrigatório")
	}

	return s.repo.DeleteEstablishment(id)
}
