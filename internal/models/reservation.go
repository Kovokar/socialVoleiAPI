package models

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	ID         int32          `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     int32          `gorm:"not null;index" json:"user_id" binding:"required"`
	ArenaID    int32          `gorm:"not null;index" json:"arena_id" binding:"required"`
	Date       time.Time      `gorm:"type:date;not null;" json:"date" binding:"required"`
	InitHour   time.Time      `gorm:"type:time;not null;" json:"init_hour" binding:"required"`
	EndHour    time.Time      `gorm:"type:time;not null;" json:"end_hour" binding:"required"`
	TotalValue float64        `gorm:"type:decimal(10,2);not null" json:"total_value" binding:"required,gt=0"` // valor em centavos
	Status     string         `gorm:"size:20;not null;default:'pendente'" json:"status"`                      // pendente, confirmada, cancelada, concluida
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`

	User  User   `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Arena Arena  `gorm:"foreignKey:ArenaID" json:"arena,omitempty"`
	Match *Match `gorm:"foreignKey:ReservationID" json:"match,omitempty"`
}
