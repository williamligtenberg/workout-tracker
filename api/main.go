package main

import (
	"log"
	"workout-tracker/api/config"
	db "workout-tracker/api/database"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db.Init(cfg)

	server := NewAPIServer(":8080")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
