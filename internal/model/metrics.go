package model

import (
	"errors"
	"strconv"
)

const (
	Counter = "counter"
	Gauge   = "gauge"
)


type Metrics struct {
	ID    string   `json:"id"`
	MType string   `json:"type"`
	Delta *int64   `json:"delta,omitempty"`
	Value *float64 `json:"value,omitempty"`
	Hash  string   `json:"hash,omitempty"`
}

type MetricsInput struct {
	Name string
	MType string
	RawValue string
}

func (m MetricsInput) Validate() error {
	switch m.MType {
	case Counter:
		delta, err := strconv.ParseInt(m.RawValue, 10, 64)
		if err != nil {
			return errors.New("invalid counter value")
		}
		if delta < 0 {
			return errors.New("counter value must be greater than 0")
		}
		return nil
	case Gauge:
		value, err := strconv.ParseFloat(m.RawValue, 64)
		if err != nil {
			return errors.New("invalid gauge value")
		}
		if value < 0 {
			return errors.New("gauge value must be greater than 0")
		}
		return nil
	}
	return errors.New("invalid metric type")
}


