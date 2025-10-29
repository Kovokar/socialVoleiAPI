package models

import (
	"time"

	"gorm.io/gorm"
)

type Establishment struct {
	ID        int32  `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"type:varchar(255)"`
	Phone     string `gorm:"type:varchar(20)"`
	CNPJ      string `gorm:"type:varchar(18);uniqueIndex;not null"`
	Latitude  float32
	Longitude float32
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Arenas []Arena `gorm:"foreignKey:EstablishmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"arenas,omitempty"`
}

type EstablishmentResponse struct {
	ID        int32   `json:"id"`
	Name      string  `json:"name" binding:"required,min=3,max=100"`
	Email     string  `json:"email" binding:"required,email"`
	Phone     string  `json:"phone" binding:"required,min=9,max=20"`
	CNPJ      string  `json:"cnpj" binding:"required,len=14"`
	Latitude  float32 `json:"latitude" binding:"gte=-90,lte=90"`
	Longitude float32 `json:"longitude" binding:"gte=-180,lte=180"`
}
