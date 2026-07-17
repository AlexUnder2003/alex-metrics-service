package service

import (
	"errors"
	"strconv"

	"github.com/Alexunder2003/alex-metrics-service/internal/model"
	"github.com/Alexunder2003/alex-metrics-service/internal/repository"
)

type MetricsService struct {
	repository repository.Storage
}

func NewMetricsService(repo repository.Storage) *MetricsService {
	return &MetricsService{repository: repo}
}

func (s *MetricsService) Update(mtype, name, value string) error {
	metrics := model.Metrics{
		ID:    name,
		MType: mtype,
	}

	switch mtype {
	case model.Gauge:
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return errors.New("invalid gauge value")
		}
		metrics.Value = &v

	case model.Counter:
		d, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return errors.New("invalid counter value")
		}
		metrics.Delta = &d

	default:
		return errors.New("invalid metric type")
	}

	s.repository.Update(metrics)
	return nil
}
