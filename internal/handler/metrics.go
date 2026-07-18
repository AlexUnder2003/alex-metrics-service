package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Alexunder2003/alex-metrics-service/internal/model"
	service "github.com/Alexunder2003/alex-metrics-service/internal/service/metric"
	"github.com/Alexunder2003/alex-metrics-service/internal/storage"
)

type Handler struct {
	svc *service.MetricsService
}

func New(storage *storage.MemStorage[model.Metrics]) *Handler {
	return &Handler{
		svc: service.NewMetricsService(storage),
	}
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	log.Println("Update request received")
	input := model.MetricsInput{
		Name:     r.PathValue("name"),
		MType:    r.PathValue("type"),
		RawValue: r.PathValue("value"),
	}

	if err := h.svc.Update(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")

	metric, err := h.svc.Get(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	response, err := json.Marshal(metric)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(response)
}