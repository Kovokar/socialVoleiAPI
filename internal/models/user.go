package models

import (
	"time"

	"gorm.io/gorm"
)

type GenderType string

const (
	Male   GenderType = "male"
	Female GenderType = "female"
	Other  GenderType = "other"
)

type User struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string         `json:"name" binding:"required,min=3,max=100"`
	Email     string         `gorm:"unique" json:"email" binding:"required,email"`
	Password  string         `json:"password"`
	Phone     string         `json:"phone" binding:"required,min=9,max=20"`
	Gender    GenderType     `gorm:"type:enum('male','female','other');default:'other'" json:"gender" binding:"oneof=male female other"`
	Latitude  float64        `json:"latitude" binding:"gte=-90,lte=90"`
	Longitude float64        `json:"longitude" binding:"gte=-180,lte=180"`
	MetricsId uint64         `json:"cd_metrics"`
	InviterId uint64         `json:"cd_inviter"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}
