package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Alexunder2003/alex-metrics-service/internal/model"
	"github.com/Alexunder2003/alex-metrics-service/internal/service"
	"github.com/Alexunder2003/alex-metrics-service/internal/storage"
)

func TestHandler_Update(t *testing.T) {
	type want struct {
		statusCode int
		response   string
	}
	tests := []struct {
		name  string
		input model.MetricsInput
		want  want
	}{
		{
			name:  "valid input",
			input: model.MetricsInput{Name: "test", MType: model.Counter, RawValue: "1"},
			want: want{
				statusCode: http.StatusOK,
				response:   "",
			},
		},
		{
			name:  "invalid input",
			input: model.MetricsInput{Name: "test", MType: model.Counter, RawValue: "invalid"},
			want: want{
				statusCode: http.StatusBadRequest,
				response:   "invalid counter value\n",
			},
		},
		{
			name:  "invalid type",
			input: model.MetricsInput{Name: "test", MType: "invalid", RawValue: "1"},
			want: want{
				statusCode: http.StatusBadRequest,
				response:   "invalid metric type\n",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			store := storage.NewMemStorage[model.Metrics]()
			router := NewRouter(Mount{
				Pattern: "/",
				Router:  MetricsRouter(New(service.NewMetricsService(store))),
			})

			url := fmt.Sprintf(
				"/update/%s/%s/%s",
				test.input.MType,
				test.input.Name,
				test.input.RawValue,
			)
			r := httptest.NewRequest(http.MethodPost, url, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, r)

			assert.Equal(t, test.want.statusCode, w.Code)
			assert.Equal(t, test.want.response, w.Body.String())
		})
	}
}

func TestHandler_Get(t *testing.T) {
	delta := int64(1)

	type want struct {
		statusCode int
		response   string
	}
	tests := []struct {
		name    string
		fixture model.Metrics
		want    want
	}{
		{
			name:    "valid input",
			fixture: model.Metrics{ID: "test", MType: model.Counter, Delta: &delta},
			want: want{
				statusCode: http.StatusOK,
				response:   `{"id":"test","type":"counter","delta":1}`,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			store := storage.NewMemStorage[model.Metrics]()
			assert.NoError(t, store.Update(test.fixture.ID, test.fixture))

			router := NewRouter(Mount{
				Pattern: "/",
				Router:  MetricsRouter(New(service.NewMetricsService(store))),
			})

			r := httptest.NewRequest(http.MethodGet, "/value/counter/test", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, r)

			assert.Equal(t, test.want.statusCode, w.Code)
			assert.Equal(t, test.want.response, w.Body.String())
		})
	}
}
