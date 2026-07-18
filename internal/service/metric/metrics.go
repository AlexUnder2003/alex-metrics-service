package service

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Alexunder2003/alex-metrics-service/internal/model"
	"github.com/Alexunder2003/alex-metrics-service/internal/repository"
)

type MetricsService struct {
	storage repository.MetricsStorage
}

func NewMetricsService(storage repository.MetricsStorage) *MetricsService {
	return &MetricsService{storage: storage}
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

		key := model.Counter + ":" + input.Name
		if current, err := s.storage.Get(key); err == nil {
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

	return s.storage.Update(metric)
}