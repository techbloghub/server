package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/techbloghub/server/internal/http/router"
)

func TestPingEndpoint(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Create a new r instance using the private setRouter function
	r := gin.Default()
	router.InitRouter(r)

	// Create a test HTTP recorder
	w := httptest.NewRecorder()

	// Create a test request to the /ping endpoint
	req, err := http.NewRequest("GET", "/ping", nil)
	assert.NoError(t, err)

	// Serve the request using the router
	r.ServeHTTP(w, req)

	// Assert the response code is 200
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert the response body is "pong"
	assert.Equal(t, "pong\n", w.Body.String())
}
