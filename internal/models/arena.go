package models

import (
	"time"

	"gorm.io/gorm"
)

type Arena struct {
	ID              uint32         `gorm:"primaryKey;autoIncrement" json:"id"`
	EstablishmentID uint32         `gorm:"not null;index" json:"establishment_id" binding:"required"`
	Name            string         `gorm:"type:varchar(255);not null" json:"name" binding:"required,min=3,max=100"`
	PlayerCapacity  int            `gorm:"not null" json:"player_capacity" binding:"required,min=1,max=100"`
	PriceHour       float64        `gorm:"type:decimal(10,2);not null" json:"price_hour" binding:"required,min=0"` // valor em centavos
	Description     string         `gorm:"type:text" json:"description"`
	Active          bool           `gorm:"default:true" json:"active"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`

	Establishment  Establishment    `gorm:"foreignKey:EstablishmentID" json:"-"`
	AvaliableTimes []AvaliableTimes `gorm:"foreignKey:ArenaID" json:"avaliable_times,omitempty"`
	Reservations   []Reservation    `gorm:"foreignKey:ArenaID" json:"reservations,omitempty"`
	Matchs         []Match          `gorm:"foreignKey:ArenaID" json:"matchs,omitempty"`
}
