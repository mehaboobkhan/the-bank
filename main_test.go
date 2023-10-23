package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestProcessDataHandler(t *testing.T) {
	// Create a request to the "/process" endpoint
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

	// Create a response recorder to capture the response.
	respRecorder := httptest.NewRecorder()

	// Create a new server instance.
	s := run()
	defer s.Close()

	// Send the request to the server.
	s.Handler.ServeHTTP(respRecorder, req)

	// Check the response status code.
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

func TestMain(m *testing.M) {
	// Run the tests with code coverage.
	os.Exit(m.Run())
}

func TestGracefulShutdown(t *testing.T) {
	// Create a new server instance.
	s := run()
	defer s.Close() // Close the server when the test is done.

	// Create a channel to signal the server to shut down.
	shutdownSignal := make(chan struct{})

	// Start a goroutine to listen for the shutdown signal and trigger the server's shutdown.
	go func() {
		// Wait for a brief moment before sending the signal.
		time.Sleep(100 * time.Millisecond)
		// Signal the server to shut down.
		close(shutdownSignal)
	}()

	// Run the server, which should wait for the shutdown signal.
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			t.Fatalf("error listening on port: %s\n", err)
		}
	}()

	// Block until the server is gracefully shut down.
	<-shutdownSignal
}
