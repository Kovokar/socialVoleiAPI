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
	ID        uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string         `gorm:"size:100;not null" json:"name" binding:"required,min=3,max=100"`
	Birthdate time.Time      `gorm:"not null" json:"birthdate" binding:"required"`
	Email     string         `gorm:"size:100;uniqueIndex;not null" json:"email" binding:"required,email"`
	Password  string         `json:"password"`
	Phone     string         `json:"phone" binding:"required,min=9,max=20"`
	Gender    GenderType     `gorm:"type:text;default:'other'" json:"gender" binding:"oneof=male female other"`
	Latitude  float64        `json:"latitude" binding:"gte=-90,lte=90"`
	Longitude float64        `json:"longitude" binding:"gte=-180,lte=180"`
	Photo     string         `json:"photo"`
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

type UsuarioLogin struct {
	Email string `json:"email" binding:"required,email"`
	Senha string `json:"senha" binding:"required,min=6"`
}

// UsuarioResponse representa a resposta sem dados sensíveis
type UsuarioResponse struct {
	ID        uint      `json:"id"`
	Nome      string    `json:"nome"`
	Idade     int       `json:"idade"`
	Email     string    `json:"email"`
	Foto      string    `json:"foto,omitempty"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	CreatedAt time.Time `json:"created_at"`
}
