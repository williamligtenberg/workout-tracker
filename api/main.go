package main

import (
	"log"
	"workout-tracker/api/auth"
	"workout-tracker/api/config"

	"os"
	db "workout-tracker/api/database"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	secret := os.Getenv("JWT_SECRET")

	auth.Init(secret)
	db.Init(cfg)

	server := NewAPIServer(":8080")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
