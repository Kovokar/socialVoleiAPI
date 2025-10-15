package models

import "time"

type Arena struct {
	ID                  uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	EstablishmentID     uint64    `gorm:"not null;index" json:"establishment_id" binding:"required"`
	Name                string    `gorm:"type:varchar(255);not null" json:"name" binding:"required,min=3,max=100"`
	CapacidadeJogadores int       `gorm:"not null" json:"capacidade_jogadores" binding:"required,min=1"`
	PriceHour           float64   `gorm:"type:decimal(10,2);not null" json:"price_hour" binding:"required,gt=0"` // valor em centavos
	Description         string    `gorm:"type:text" json:"description"`
	Active              bool      `gorm:"default:true" json:"active"`
	CreatedAt           time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt           time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relacionamento reverso (belongs to)
	Establishment Establishment `gorm:"foreignKey:EstablishmentID" json:"-"`
}
