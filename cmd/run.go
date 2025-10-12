package cmd

import (
	"socialVoleiAPI/internal/database"
	"socialVoleiAPI/internal/server"
)

func Run() {

	database.StartDb()

	server := server.NewServer()
	server.Run()
}
