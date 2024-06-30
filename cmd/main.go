package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/asdavies/internal/config"
	"github.com/asdavies/internal/routing"
	"github.com/asdavies/public"
)

func main() {
    // Initialize the configuration
    appConfig := config.InitConfig()

    public.RegisterGlobalStyles(appConfig.GlobalCSS...)

	router := routing.NewRouter()

	log.Printf("Starting server on :%d\n", appConfig.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", appConfig.Port), router); err != nil {
		log.Fatalf("Could not start server: %s", err)
	}
}
