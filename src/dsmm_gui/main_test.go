package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestDsmmForm(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DsmmForm)

	handler.ServeHTTP(rr, req)

	// Test server response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("dsmmRatings handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Test if html template is returned properly
	if !strings.Contains(rr.Body.String(), "<label>Preservability</label>") {
		t.Errorf("dsmmRatings handler returned unexpected body")
	}

	// Test if form entries are processed correctly
	initialFormItems := map[string]string{
		"Preservability":               "adhoc",
		"Accessibility":                "minimal",
		"Usability":                    "intermediate",
		"ProductionSustainability":     "advanced",
		"DataQualityAssurance":         "optimal",
		"DataQualityControlMonitoring": "adhoc",
		"DataQualityAssessment":        "minimal",
		"TransparencyTraceability":     "minimal",
		"DataIntegrity":                "optimal",
	}

	req.Form = make(url.Values)
	for k, v := range initialFormItems {
		req.Form.Add(k, v)
	}

	wantForm := url.Values{
		"Preservability":               []string{"adhoc"},
		"Accessibility":                []string{"minimal"},
		"Usability":                    []string{"intermediate"},
		"ProductionSustainability":     []string{"advanced"},
		"DataQualityAssurance":         []string{"optimal"},
		"DataQualityControlMonitoring": []string{"adhoc"},
		"DataQualityAssessment":        []string{"minimal"},
		"TransparencyTraceability":     []string{"minimal"},
		"DataIntegrity":                []string{"optimal"},
	}
	if !reflect.DeepEqual(req.Form, wantForm) {
		t.Fatalf("req.Form = %v, want %v", req.Form, wantForm)
	}
}

func TestDsmmResults(t *testing.T) {
	req, err := http.NewRequest("GET", "/dsmm_results", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DsmmResults)

	handler.ServeHTTP(rr, req)

	// Test server response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("DsmmResults handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Test if html template is returned properly
	if !strings.Contains(rr.Body.String(), "<h1><strong>Success!</strong></h1>") {
		t.Errorf("dsmmRatings handler returned unexpected body")
	}
}

func TestDsmmWriteSnippetToBrowser(t *testing.T) {
	req, err := http.NewRequest("GET", "/dsmm_xml", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DsmmWriteSnippetToBrowser)

	handler.ServeHTTP(rr, req)

	// Test server response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("DsmmResults handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Test if html template is returned properly
	if !strings.Contains(rr.Body.String(), "<gco:CharacterString>Data Stewardship Maturity Assessment</gco:CharacterString>") {
		t.Errorf("dsmmRatings handler returned unexpected body")
	}
}
