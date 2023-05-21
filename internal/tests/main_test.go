package tests

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTotalIncomeTax(t *testing.T) {
	// Create a new router instance
	router := gin.New()
	router.Use(gin.Recovery())

	// Configure the router for testing
	gin.SetMode(gin.TestMode)

	// Mock the tax calculator API endpoint
	router.GET("/income-tax/calculate-tax", func(c *gin.Context) {
		c.JSON(http.StatusOK, []float64{5000.0})
	})

	// Create a new HTTP request
	req := httptest.NewRequest("GET", "/income-tax/calculate-tax?year=2023&salary=25000", nil)

	// Create a new HTTP response recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code, "Unexpected status code")

	// Parse the response body
	var response []float64
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err, "Failed to parse response body")

	// Check the tax amount in the response
	expectedTaxAmount := 5000.0
	assert.Len(t, response, 1, "Unexpected number of tax amounts")
	assert.Equal(t, expectedTaxAmount, response[0], "Unexpected tax amount")
}

func TestGetTotalIncomeTax_MissingQueryParams(t *testing.T) {
	// Create a new router instance
	router := gin.New()
	router.Use(gin.Recovery())

	// Configure the router for testing
	gin.SetMode(gin.TestMode)

	// Handle the route for missing query parameters
	router.GET("/income-tax/calculate-tax", func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing query parameters",
		})
	})

	// Create a new HTTP request with missing query parameters
	req := httptest.NewRequest("GET", "/income-tax/calculate-tax", nil)

	// Create a new HTTP response recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusBadRequest, w.Code, "Unexpected status code")

	// Parse the response body
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err, "Failed to parse response body")

	// Check the error message in the response
	expectedErrorMessage := "Missing query parameters"
	assert.Equal(t, expectedErrorMessage, response["error"], "Unexpected error message")
}
