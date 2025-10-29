package repositories

import (
	"socialVoleiAPI/internal/models"

	"gorm.io/gorm"
)

type EstablishmentRepository struct {
	db *gorm.DB
}

func NewEstablishmentRepository(db *gorm.DB) *EstablishmentRepository {
	return &EstablishmentRepository{db: db}
}

func (r *EstablishmentRepository) CreateEstablishment(Establishment *models.Establishment) error {
	return r.db.Create(Establishment).Error
}

func (r *EstablishmentRepository) FindAllEstablishments() ([]models.Establishment, error) {
	var Establishments []models.Establishment
	err := r.db.Find(&Establishments).Error
	return Establishments, err
}

func (r *EstablishmentRepository) FindEstablishmentByID(Establishment models.Establishment, intId int) (models.Establishment, error) {
	err := r.db.First(&Establishment, intId).Error
	return Establishment, err
}
