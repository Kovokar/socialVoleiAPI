package models

import (
	"time"

	"gorm.io/gorm"
	// "gorm.io/gorm"
)

type GenderType string

const (
	Male   GenderType = "male"
	Female GenderType = "female"
	Other  GenderType = "other"
)

type User struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement"`
	Name      string         `gorm:"size:100;not null" binding:"required,min=3,max=100"`
	Birthdate time.Time      `gorm:"not null" binding:"required"`
	Email     string         `gorm:"size:100;uniqueIndex;not null" binding:"required,email"`
	Password  string         `gorm:"not null" binding:"required"`
	Phone     string         `gorm:"size:50" binding:"required,min=9,max=20"`
	Gender    GenderType     `gorm:"type:text;default:'other'" binding:"oneof=male female other"`
	Latitude  float64        `json:"latitude" binding:"gte=-90,lte=90,required"`
	Longitude float64        `json:"longitude" binding:"gte=-180,lte=180, required"`
	Photo     string         `gorm:"default:'null'" json:"photo"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	CreatedsMatchs []Match              `gorm:"foreignKey:CreaterID" json:"createds_matchs,omitempty"`
	Participations []MatchParticipation `gorm:"foreignKey:UserID" json:"participations,omitempty"`
	Reservation    []Reservation        `gorm:"foreignKey:UserID" json:"reservations,omitempty"`
	// AvaliacoesFeitas    []Avaliacao          `gorm:"foreignKey:AvaliadorID" json:"avaliacoes_feitas,omitempty"`
	// AvaliacoesRecebidas []Avaliacao          `gorm:"foreignKey:AvaliadoID" json:"avaliacoes_recebidas,omitempty"`
	// TimesCriados        []Time               `gorm:"foreignKey:CriadorID" json:"times_criados,omitempty"`
	// 	MembrosTime         []MembroTime         `gorm:"foreignKey:UsuarioID" json:"membros_time,omitempty"`
}

type UserLogin struct {
	Email string `json:"email" binding:"required,email"`
	Senha string `json:"senha" binding:"required,min=6"`
}

// UsuarioResponse representa a resposta sem dados sensíveis
type UserResponse struct {
	ID        uint64     `json:"id"`
	Name      string     `json:"name"`
	Birthdate time.Time  `json:"birthdate" `
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	Gender    GenderType `json:"gender"`
	Latitude  float64    `json:"latitude"`
	Longitude float64    `json:"longitude"`
	Photo     string     `json:"photo"`
}
