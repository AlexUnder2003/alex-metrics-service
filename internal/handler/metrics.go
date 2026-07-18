package handler

import (
	"net/http"

	"github.com/Alexunder2003/alex-metrics-service/internal/model"
	"github.com/Alexunder2003/alex-metrics-service/internal/repository"
	service "github.com/Alexunder2003/alex-metrics-service/internal/service/metric"
)

type Handler struct {
	svc *service.MetricsService
}

func New(storage repository.MetricsStorage) *Handler {
	return &Handler{
		svc: service.NewMetricsService(storage),
	}
}


func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	input := model.MetricsInput{
		Name: r.PathValue("name"),
		MType: r.PathValue("type"),
		RawValue: r.PathValue("value"),
	}

	if err := h.svc.Update(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
