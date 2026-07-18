package app

import (
	"net/http"

	"github.com/Alexunder2003/alex-metrics-service/internal/config"
	"github.com/Alexunder2003/alex-metrics-service/internal/handler"
	"github.com/Alexunder2003/alex-metrics-service/internal/model"
	"github.com/Alexunder2003/alex-metrics-service/internal/storage"
)

type App struct {
	cfg config.Config
	mux *http.ServeMux
}

func NewApp(cfg config.Config) *App {
	storage := storage.NewMemStorage[model.Metrics]()
	handler := handler.New(storage)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /update/{type}/{name}/{value}", handler.Update)
	mux.HandleFunc("GET /value/{type}/{name}", handler.Get)

	return &App{cfg: cfg, mux: mux}
}

func (a *App) Run() error {
	return http.ListenAndServe(a.cfg.Address, a.mux)
}