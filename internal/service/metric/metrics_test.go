package service

import (
	"testing"

	"github.com/Alexunder2003/alex-metrics-service/internal/model"
	"github.com/Alexunder2003/alex-metrics-service/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestMetricsService_Update(t *testing.T) {
	tests := []struct {
		name    string
		input   model.MetricsInput
		wantErr string
	}{
		{
			name:  "valid input",
			input: model.MetricsInput{Name: "test", MType: model.Counter, RawValue: "1"},
		},
		{
			name:    "invalid value",
			input:   model.MetricsInput{Name: "test", MType: model.Counter, RawValue: "invalid"},
			wantErr: "invalid counter value",
		},
		{
			name:    "invalid type",
			input:   model.MetricsInput{Name: "test", MType: "invalid", RawValue: "1"},
			wantErr: "invalid metric type",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			storage := storage.NewMemStorage[model.Metrics]()
			svc := NewMetricsService(storage)

			err := svc.Update(test.input)

			if test.wantErr == "" {
				assert.NoError(t, err)
				return
			}

			assert.EqualError(t, err, test.wantErr)
		})
	}
}
