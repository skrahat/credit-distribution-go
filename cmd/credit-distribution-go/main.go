package main

import (
	"log"
	"net/http"

	"credit-distribution-go/internal/config"
	"credit-distribution-go/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {
    cfg := config.LoadConfig()

    r := mux.NewRouter()
    handlers.RegisterHandlers(r)

    log.Printf("Starting server on %s", cfg.Server.Address)
    if err := http.ListenAndServe(cfg.Server.Address, r); err != nil {
        log.Fatalf("could not start server: %v\n", err)
    }
}
