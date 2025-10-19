package migrations

import (
	"socialVoleiAPI/internal/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Arena{})
	db.AutoMigrate(models.Establishment{})
	db.AutoMigrate(models.AvaliableTimes{})
	db.AutoMigrate(models.Match{})
	db.AutoMigrate(models.Reservation{})
	db.AutoMigrate(models.MatchParticipation{})
}
