package main

import (
	"log"
	"net/http"

	"github.com/Alexunder2003/alex-metrics-service/internal/config"
	"github.com/Alexunder2003/alex-metrics-service/internal/handler"
	"github.com/Alexunder2003/alex-metrics-service/internal/repository"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	repo := repository.NewMemStorage()
	metricsHandler := handler.New(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /update/{type}/{name}/{value}", metricsHandler.Update)

	if err := http.ListenAndServe(cfg.Address, mux); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
