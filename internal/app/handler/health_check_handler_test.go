package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/vinsensiussatya/bego-training/internal/app/service/mocks"
	"github.com/vinsensiussatya/bego-training/internal/pkg/response"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler_Readiness(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		wantStatus int
		mockFunc   func(mockHcs *mocks.IHealthCheckService)
	}{
		{
			name: "OK",
			mockFunc: func(mockHcs *mocks.IHealthCheckService) {
				mockHcs.On("Ping", context.Background()).Return(response.PingResponse{
					Database: "OK",
				})
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "Error",
			mockFunc: func(mockHcs *mocks.IHealthCheckService) {
				mockHcs.On("Ping", context.Background()).Return(response.PingResponse{
					Database: "ERROR",
				})
			},
			wantStatus: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hcs := mocks.NewIHealthCheckService(t)
			tt.mockFunc(hcs)
			h := NewHealthCheckHandler(hcs)

			w := httptest.NewRecorder()
			r, err := http.NewRequest(http.MethodGet, "", nil)
			if err != nil {
				t.Fatal(err)
			}
			MyHandler(h.Readiness).ServeHTTP(w, r)

			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

func TestHealthCheckHandler_Liveness(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		wantStatus int
	}{
		{
			name:       "OK",
			wantStatus: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hcs := mocks.NewIHealthCheckService(t)
			h := NewHealthCheckHandler(hcs)

			w := httptest.NewRecorder()
			r, err := http.NewRequest(http.MethodGet, "", nil)
			if err != nil {
				t.Fatal(err)
			}
			MyHandler(h.Liveness).ServeHTTP(w, r)

			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}
