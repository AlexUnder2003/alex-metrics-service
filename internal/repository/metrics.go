package repository

import (
	"errors"

	"github.com/Alexunder2003/alex-metrics-service/internal/model"
)

type MetricsStorage struct {
	storage map[string]model.Metrics
}

func NewMetricsStorage() *MetricsStorage {
	return &MetricsStorage{
		storage: make(map[string]model.Metrics),
	}
}

func (s *MetricsStorage) Get(key string) (model.Metrics, error) {
	metric, ok := s.storage[key]
	if !ok {
		return model.Metrics{}, errors.New("metric not found")
	}
	return metric, nil
}

func (s *MetricsStorage) Update(metric model.Metrics) error {
	key := metric.MType + ":" + metric.ID

	s.storage[key] = metric

	return nil
}
