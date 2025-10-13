package models

import (
	"time"

	"gorm.io/gorm"
)

type StatusType string

const (
	Pending StatusType = "ending"
	Used    StatusType = "used"
	Expired StatusType = "expired"
)

type Invite struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint64         `gorm:"not null" json:"cd_user"`
	User      User           `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Invite    string         `json:"invite" binding:"required,min=3,max=15"`
	Status    StatusType     `gorm:"type:enum('pending','used','expired');default:pending" binding:"oneof=pending used expired"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}
