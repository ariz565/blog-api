package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// TestGetPosts tests the /posts endpoint
func TestGetPosts(t *testing.T) {
	// Set Gin to test mode to avoid noisy logging
	gin.SetMode(gin.TestMode)

	// Set up a test router and register only the route we want to test
	router := gin.Default()
	router.GET("/posts", getPosts)

	// Create a test HTTP request (GET /posts)
	req, err := http.NewRequest(http.MethodGet, "/posts", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v", err)
	}

	// Record the response using httptest
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	// Check for 200 OK
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", recorder.Code)
	}

	// Optionally: check response body (if you want)
	expected := `[` // response is a JSON array
	if recorder.Body.String()[:1] != expected {
		t.Errorf("Expected response body to start with %q, got %q", expected, recorder.Body.String())
	}
}
