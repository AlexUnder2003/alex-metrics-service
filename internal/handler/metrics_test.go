package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Alexunder2003/alex-metrics-service/internal/model"
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
			storage := storage.NewMemStorage[model.Metrics]()
			h := New(storage)

			r := httptest.NewRequest(http.MethodPost, "/update", nil)
			r.SetPathValue("type", test.input.MType)
			r.SetPathValue("name", test.input.Name)
			r.SetPathValue("value", test.input.RawValue)

			w := httptest.NewRecorder()
			h.Update(w, r)

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
			storage := storage.NewMemStorage[model.Metrics]()
			assert.NoError(t, storage.Update(test.fixture.ID, test.fixture))

			h := New(storage)

			r := httptest.NewRequest(http.MethodGet, "/value", nil)
			r.SetPathValue("name", "test")

			w := httptest.NewRecorder()
			h.Get(w, r)

			assert.Equal(t, test.want.statusCode, w.Code)
			assert.Equal(t, test.want.response, w.Body.String())
		})
	}
}
