package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/asdavies/auth/internal/routing"
)

func main() {
	port := 8080

	router := routing.NewRouter()

	log.Printf("Starting server on :%d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		log.Fatalf("Could not start server: %s", err)
	}
}
