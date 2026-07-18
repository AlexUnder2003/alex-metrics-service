package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Alexunder2003/alex-metrics-service/internal/model"
	"github.com/Alexunder2003/alex-metrics-service/internal/service"
)

type Handler struct {
	svc *service.MetricsService
}

func New(svc *service.MetricsService) *Handler {
	return &Handler{svc: svc}
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