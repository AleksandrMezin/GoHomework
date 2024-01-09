package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetNewsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/news?count=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getNewsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
