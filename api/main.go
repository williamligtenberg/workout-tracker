package main

import (
	"log"
)

func main() {
	databaseConnection()

	server := NewAPIServer(":8080")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
