package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProcessDataApproved(t *testing.T) {

	requestBody := []byte(`{
		  "income": 182428,
		  "number_of_credit_cards": 3,
		  "age": 18,
		  "politically_exposed": false,
		  "job_industry_code": "2-930 - Exterior Plants",
		  "phone_number": "886-356-0377"
		}`)

	req, err := http.NewRequest(http.MethodPost, "/process", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	respRecorder := httptest.NewRecorder()
	ProcessData(respRecorder, req)

	if respRecorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, respRecorder.Code)
	}

	var jsonResponse map[string]interface{}
	if err := json.NewDecoder(respRecorder.Body).Decode(&jsonResponse); err != nil {
		t.Fatal(err)
	}

	expectedStatus := "approved"
	if jsonResponse["status"] != expectedStatus {
		t.Errorf("Expected status %s but got %s", expectedStatus, jsonResponse["status"])
	}
}

func TestProcessDataDeclined(t *testing.T) {

	requestBody := []byte(`{
		  "income": 82428,
		  "number_of_credit_cards": 3,
		  "age": 18,
		  "politically_exposed": false,
		  "job_industry_code": "2-930 - Exterior Plants",
		  "phone_number": "886-356-0377"
		}`)

	req, err := http.NewRequest(http.MethodPost, "/process", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	respRecorder := httptest.NewRecorder()
	ProcessData(respRecorder, req)

	if respRecorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, respRecorder.Code)
	}

	var jsonResponse map[string]interface{}
	if err := json.NewDecoder(respRecorder.Body).Decode(&jsonResponse); err != nil {
		t.Fatal(err)
	}

	expectedStatus := "declined"
	if jsonResponse["status"] != expectedStatus {
		t.Errorf("Expected status %s but got %s", expectedStatus, jsonResponse["status"])
	}
}
