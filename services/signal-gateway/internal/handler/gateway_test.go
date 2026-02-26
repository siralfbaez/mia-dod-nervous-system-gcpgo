package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOrderHandler(t *testing.T) {
	req := httptest.NewRequest("POST", "/v1/order", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(OrderHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}