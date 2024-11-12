package main

import "log"

func main() {
	server := NewAPIServer(":8080")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
