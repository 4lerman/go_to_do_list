package main

import (
	"log"

	"github.com/4lerman/go_to_do_list/cmd/to_do_list/api"
	"github.com/4lerman/go_to_do_list/internal/db"
)

func main() {
	db := db.NewSyncMapStorage()

	server := api.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal("Error when running server", err)
	}
}
