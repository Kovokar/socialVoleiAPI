package repositories

import (
	"fmt"
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

func (r *EstablishmentRepository) BulkCreateEstablishment(Establishments *[]models.Establishment) error {
	return r.db.Create(Establishments).Error
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

func (r *EstablishmentRepository) UpdateEstablishment(id string, data models.Establishment) error {
	return r.db.Model(&models.Establishment{}).
		Where("id = ?", id).
		Updates(data).Error
}

func (r *EstablishmentRepository) DeleteEstablishment(id string) error {
	result := r.db.Delete(&models.Establishment{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("estabelecimento não encontrado")
	}

	return nil
}
