package main

import (
  "net/http"
  "net/http/httptest"
  "testing"
)

var routes = []string{
  "/index.html",
  "/dsmm.html",
  // "/poop.html",
}

func TestServeTemplate(t *testing.T) {
  for _, route := range routes {
    req, err := http.NewRequest("GET", route, nil)
    if err != nil {
      t.Fatal(err)
    }
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(serveTemplate)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
      t.Errorf("handler for route %v returned wrong status code: got %v, want %v", route, status, http.StatusOK)
    }
  }
}
