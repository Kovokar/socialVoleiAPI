package models

import (
	"time"

	"gorm.io/gorm"
)

type MatchStatusType string

// const (
// 	Pending MatchStatusType = "ending"
// 	Used    MatchStatusType = "used"
// 	Expired MatchStatusType = "expired"
// )

type Match struct {
	ID               int32          `gorm:"primaryKey;autoIncrement" json:"id"`
	CreaterID        int32          `gorm:"not null;index" json:"creater_id" binding:"required"`
	ArenaID          int32          `gorm:"not null;index" json:"arena_id" binding:"required"`
	ReservationID    uint           `gorm:"uniqueIndex" json:"reservation_id"`
	Type             string         `gorm:"size:20;not null" json:"type" binding:"required,oneof=public private"` // publica, privada
	Date             time.Time      `gorm:"type:date;not null;" json:"date" binding:"required"`
	InitHour         time.Time      `gorm:"type:time;not null;" json:"init_hour" binding:"required"`
	EndHour          time.Time      `gorm:"type:time;not null;" json:"end_hour" binding:"required"`
	MaxPlayeCapacity int            `gorm:"not null" json:"max_player_capacity" binding:"required,min=1"`
	InviteCode       string         `gorm:"size:15;uniqueIndex" json:"invite_code,omitempty"`
	Status           string         `gorm:"size:20;not null;default:'aguardando'" json:"status"` // aguardando, em_andamento, finalizada, cancelada
	CreatedAt        time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`

	Creater        User                 `gorm:"foreignKey:CreaterID" json:"creater,omitempty"`
	Arena          Arena                `gorm:"foreignKey:ArenaID" json:"arena,omitempty"`
	Reservation    Reservation          `gorm:"foreignKey:ReservationID" json:"reservation,omitempty"`
	Participations []MatchParticipation `gorm:"foreignKey:MatchID" json:"participations,omitempty"`
}
