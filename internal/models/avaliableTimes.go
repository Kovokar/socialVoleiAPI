package models

import (
	"time"

	"gorm.io/gorm"
)

type AvaliableTimes struct {
	ID        int32     `gorm:"primaryKey;autoIncrement" json:"id"`
	ArenaID   int32     `gorm:"not null;index" json:"arena_id" binding:"required"`
	Date      time.Time `gorm:"type:date;not null;" json:"date" binding:"required"`
	InitHour  time.Time `gorm:"type:time;not null;" json:"init_hour" binding:"required"`
	EndHour   time.Time `gorm:"type:time;not null;" json:"end_hour" binding:"required"`
	Available bool      `gorm:"default:true" json:"available"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Arena Arena `gorm:"foreignKey:ArenaID" json:"-"`
}

func (AvaliableTimes) TableName() string {
	return "reservation"
}
