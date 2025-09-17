package database

import (
	"log"
	"socialVoleiAPI/internal/database/migrations"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDb() {
	strConn := "host=localhost user=admin password=123456 dbname=books port=25431 sslmode=disable"

	conn, err := gorm.Open(postgres.Open(strConn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	db = conn
	config, _ := db.DB()

	config.SetConnMaxLifetime(10)
	config.SetMaxOpenConns(100)
	config.SetMaxIdleConns(10)

	migrations.RunMigrations(db)
}

func GetDatabase() *gorm.DB {
	return db
}
