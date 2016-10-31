package main

import (
	"net/http"
  "net/http/httptest"
  "strings"
	"testing"
)

func TestDsmmRatings(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(dsmmRatings)

  handler.ServeHTTP(rr, req)

  if status := rr.Code; status != http.StatusOK {
    t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
  }

  if !strings.Contains(rr.Body.String(), "<label>Preservability</label>") {
    t.Errorf("handler returned unexpected body")
  }
}
