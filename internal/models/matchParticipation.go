package models

import (
	"time"

	"gorm.io/gorm"
)

type MatchParticipationStatusType string

// const (
// 	Pending StatusType = "ending"
// 	Used    StatusType = "used"
// 	Expired StatusType = "expired"
// )

type MatchParticipation struct {
	ID        int32          `gorm:"primaryKey;autoIncrement" json:"id"`
	MatchID   int32          `gorm:"not null;index" json:"match_id" binding:"required"`
	UserID    int32          `gorm:"not null;index" json:"creater_id" binding:"required"`
	Status    string         `gorm:"size:20;not null;default:'pendente'" json:"status"` // pendente, aceito, recusado
	Winner    bool           `gorm:"default:false" json:"winner"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Match Match `gorm:"foreignKey:MatchID" json:"match,omitempty"`
	User  User  `gorm:"foreignKey:UserID" json:"user,omitempty"`
}
