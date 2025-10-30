package models

import (
	"time"

	"gorm.io/gorm"
)

type Establishment struct {
	ID        int32          `gorm:"primaryKey;autoIncrement"`
	Name      string         `gorm:"type:varchar(255);not null" binding:"required,min=3,max=100"`
	Email     string         `gorm:"type:varchar(255)" binding:"required,email"`
	Phone     string         `gorm:"type:varchar(20)" binding:"required,min=9,max=20"`
	CNPJ      string         `gorm:"type:varchar(18);uniqueIndex;not null" binding:"required,len=14"`
	Latitude  float32        `binding:"gte=-90,lte=90"`
	Longitude float32        `binding:"gte=-180,lte=180"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Arenas []Arena `gorm:"foreignKey:EstablishmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"arenas,omitempty"`
}

type EstablishmentResponse struct {
	ID        int32   `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Phone     string  `json:"phone"`
	CNPJ      string  `json:"cnpj"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
