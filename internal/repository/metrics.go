package repository

import "github.com/Alexunder2003/alex-metrics-service/internal/model"

type Storage interface {
	Update(metrics model.Metrics)
}

type MemStorage struct {
	storage map[string]model.Metrics
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		storage: make(map[string]model.Metrics),
	}
}

func (m *MemStorage) Update(metrics model.Metrics) {
	m.storage[metrics.ID] = metrics
}
