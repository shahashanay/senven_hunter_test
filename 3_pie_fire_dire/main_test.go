package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBeefSummaryHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/beef/summary", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(beefSummaryHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Response is not valid JSON: %v", err)
	}

	if _, ok := response["beef"]; !ok {
		t.Errorf("Response JSON does not contain 'beef' key")
	}
}
