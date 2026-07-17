package handler

import (
	"net/http"

	"github.com/Alexunder2003/alex-metrics-service/internal/repository"
	"github.com/Alexunder2003/alex-metrics-service/internal/service"
)

type Handler struct {
	repository repository.Storage
	svc *service.MetricsService
}

func New(repo repository.Storage) *Handler {
	svc := service.NewMetricsService(repo)
	return &Handler{repository: repo, svc: svc}
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	mtype := r.PathValue("type")
	name := r.PathValue("name")
	value := r.PathValue("value")

	if name == "" || mtype == "" || value == "" {
		http.Error(w, "type, name and value are required", http.StatusNotFound)
		return
	}

	if err := h.svc.Update(mtype, name, value); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
