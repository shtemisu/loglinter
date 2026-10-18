package a

import (
	"errors"
	"log"
)

func testEnglish() {
	err := errors.New("connection to db")
	port := "8080"
	log.Fatal("failed to connect server")
	log.Println("failed to fetch data")
	log.Printf("failed to fetch data: %w", err)
	log.Fatalf("failed to connect server: %s", port)
}
