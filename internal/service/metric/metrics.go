package service

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Alexunder2003/alex-metrics-service/internal/model"
	"github.com/Alexunder2003/alex-metrics-service/internal/repository"
	"github.com/Alexunder2003/alex-metrics-service/internal/storage"
)

type MetricsService struct {
	repository *repository.MetricsRepository
}

func NewMetricsService(storage *storage.MemStorage[model.Metrics]) *MetricsService {
	repository := repository.NewMetricsRepository(storage)
	return &MetricsService{repository: repository}
}

func (s *MetricsService) Update(input model.MetricsInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	metric := model.Metrics{
		ID:    input.Name,
		MType: input.MType,
	}

	switch input.MType {
	case model.Counter:
		delta, err := strconv.ParseInt(input.RawValue, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid counter value: %w", err)
		}

		key := input.Name
		if current, err := s.repository.Get(key); err == nil {
			delta += *current.Delta
		}
		metric.Delta = &delta

	case model.Gauge:
		value, err := strconv.ParseFloat(input.RawValue, 64)
		if err != nil {
			return fmt.Errorf("invalid gauge value: %w", err)
		}
		metric.Value = &value

	default:
		return errors.New("invalid metric type")
	}

	return s.repository.Update(metric)
}


func (s *MetricsService) Get(id string) (model.Metrics, error) {
	metric, err := s.repository.Get(id)
	if err != nil {
		return model.Metrics{}, errors.New("metric not found")
	}
	return metric, nil
}