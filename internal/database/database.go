package database

import (
	"log"
	"socialVoleiAPI/internal/database/migrations"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDb() {
	strConn :=
		`host=localhost 
	user=app_user 
	password=app_password 
	dbname=app_db 
	port=6543 
	sslmode=disable`

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
