package service

import (
	"github.com/Alexunder2003/alex-metrics-service/internal/model"
)

type MetricsStorage interface {
	Get(string) (model.Metrics, error)
	Update(model.Metrics) error
}
