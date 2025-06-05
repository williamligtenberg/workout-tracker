package main

import (
	"log"

	"github.com/williamligtenberg/workout-tracker/auth"
	"github.com/williamligtenberg/workout-tracker/config"

	"os"

	db "github.com/williamligtenberg/workout-tracker/database"
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
