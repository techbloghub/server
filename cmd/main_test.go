package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingEndpoint(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Create a new router instance using the private setRouter function
	router := SetRouter()

	// Create a test HTTP recorder
	w := httptest.NewRecorder()

	// Create a test request to the /ping endpoint
	req, err := http.NewRequest("GET", "/ping", nil)
	assert.NoError(t, err)

	// Serve the request using the router
	router.ServeHTTP(w, req)

	// Assert the response code is 200
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert the response body is "pong"
	assert.Equal(t, "pong", w.Body.String())
}
