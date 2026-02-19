package go_auth_api_tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerRegister(t *testing.T) {
	tests := []struct {
		name       string
		body       string
		method     string
		wantStatus int
	}{
		{
			name:       "201(valid request",
			body:       `{"email":"test@test.com","password":"123456"}`,
			method:     http.MethodPost,
			wantStatus: http.StatusCreated,
		},
		{
			name:       "405(method not allowed)",
			body:       `{"email":"test@test.com","password":"123456"}`,
			method:     http.MethodGet,
			wantStatus: http.StatusMethodNotAllowed,
		},
		{
			name:       "400(empty json)",
			body:       `{}`,
			method:     http.MethodPost,
			wantStatus: http.StatusBadRequest,
		}, {
			name:       "400(invalid email)",
			body:       `{"email":"testtest.com","password":"123456"}`,
			method:     http.MethodPost,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "400(invalid password)",
			body:       `{"email":"test@test.com","password":"123"}`,
			method:     http.MethodPost,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "400(invalid json)",
			body:       `{"email":"test@test.com"}`,
			method:     http.MethodPost,
			wantStatus: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/register", bytes.NewBufferString(tt.body))
			rep := httptest.NewRecorder()
			HandlerRegister(rep, req)
			if rep.Code != tt.wantStatus {
				t.Errorf("want status %d, got %d", tt.wantStatus, rep.Code)
			}
		})
	}
}
