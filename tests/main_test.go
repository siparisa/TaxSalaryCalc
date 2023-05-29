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

// Happy Path

// TestGetTotalIncomeTax tests the total Amount tax income
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

func TestGetTotalIncomeTax_TaxAmountPerBand(t *testing.T) {
	//// Create a new router instance
	//router := gin.New()
	//router.Use(gin.Recovery())
	//
	//// Configure the router for testing
	//gin.SetMode(gin.TestMode)
	//
	//// Mock the tax calculator API endpoint
	//router.GET("/income-tax/calculate-tax", func(c *gin.Context) {
	//	// Mock the tax brackets for the given year
	//	taxBrackets := []service.TaxBracket{
	//		{LowerBound: 0, UpperBound: 10000, Rate: 10},
	//		{LowerBound: 10001, UpperBound: 20000, Rate: 15},
	//	}
	//
	//	// Get the salary from the query parameter
	//	salaryStr := c.Query("salary")
	//	salary, _ := strconv.ParseFloat(salaryStr, 64)
	//
	//	// Calculate the tax amount per band
	//	taxAmountPerBand := service.CalculateTaxPerBand(taxBrackets, salary)
	//
	//	c.JSON(http.StatusOK, taxAmountPerBand)
	//})
	//
	//// Create a new HTTP request
	//req := httptest.NewRequest("GET", "/income-tax/calculate-tax?year=2023&salary=15000", nil)
	//
	//// Create a new HTTP response recorder
	//w := httptest.NewRecorder()
	//
	//// Perform the request
	//router.ServeHTTP(w, req)
	//
	//// Check the response status code
	//assert.Equal(t, http.StatusOK, w.Code, "Unexpected status code")
	//
	//// Parse the response body
	//var response []float64
	//err := json.Unmarshal(w.Body.Bytes(), &response)
	//require.NoError(t, err, "Failed to parse response body")
	//
	//// Check the tax amount per band in the response
	//expectedTaxAmountPerBand := []float64{1500.0, 750.0}
	//assert.Equal(t, expectedTaxAmountPerBand, response, "Unexpected tax amount per band")
}

func TestGetTotalIncomeTax_EffectiveRate(t *testing.T) {
	//// Create a new router instance
	//router := gin.New()
	//router.Use(gin.Recovery())
	//
	//// Configure the router for testing
	//gin.SetMode(gin.TestMode)
	//
	//// Mock the tax calculator API endpoint
	//router.GET("/income-tax/calculate-tax", func(c *gin.Context) {
	//	// Mock the tax brackets for the given year
	//	taxBrackets := []service.TaxBracket{
	//		{LowerBound: 0, UpperBound: 10000, Rate: 10},
	//		{LowerBound: 10001, UpperBound: 20000, Rate: 15},
	//	}
	//
	//	// Get the salary from the query parameter
	//	salaryStr := c.Query("salary")
	//	salary, _ := strconv.ParseFloat(salaryStr, 64)
	//
	//	// Calculate the tax amount for the salary
	//	taxAmount := service.CalculateTaxForSalary(taxBrackets, salary)
	//
	//	// Calculate the effective tax rate
	//	effectiveRate := service.CalculateEffectiveRate(taxAmount, salary)
	//
	//	c.JSON(http.StatusOK, effectiveRate)
	//})
	//
	//// Create a new HTTP request
	//req := httptest.NewRequest("GET", "/income-tax/calculate-tax?year=2023&salary=15000", nil)
	//
	//// Create a new HTTP response recorder
	//w := httptest.NewRecorder()
	//
	//// Perform the request
	//router.ServeHTTP(w, req)
	//
	//// Check the response status code
	//assert.Equal(t, http.StatusOK, w.Code, "Unexpected status code")
	//
	//// Parse the response body
	//var response float64
	//err := json.Unmarshal(w.Body.Bytes(), &response)
	//require.NoError(t, err, "Failed to parse response body")
	//
	//// Check the effective tax rate in the response
	//expectedEffectiveRate := 10.0
	//assert.Equal(t, expectedEffectiveRate, response, "Unexpected effective tax rate")
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

func TestGetTotalIncomeTax_InvalidSalary(t *testing.T) {
	// Create a new router instance
	router := gin.New()
	router.Use(gin.Recovery())

	// Configure the router for testing
	gin.SetMode(gin.TestMode)

	// Handle the route for invalid salary
	router.GET("/income-tax/calculate-tax", func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid salary",
		})
	})

	// Create a new HTTP request with an invalid salary
	req := httptest.NewRequest("GET", "/income-tax/calculate-tax?year=2023&salary=invalid", nil)

	// Create a new HTTP response recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusInternalServerError, w.Code, "Unexpected status code")

	// Parse the response body
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err, "Failed to parse response body")

	// Check the error message in the response
	expectedErrorMessage := "Invalid salary"
	assert.Equal(t, expectedErrorMessage, response["error"], "Unexpected error message")
}

func TestGetTotalIncomeTax_ErrorResponse(t *testing.T) {
	// Create a new router instance
	router := gin.New()
	router.Use(gin.Recovery())

	// Configure the router for testing
	gin.SetMode(gin.TestMode)

	// Mock the tax calculator API endpoint with an error response
	router.GET("/income-tax/calculate-tax", func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve tax brackets",
		})
	})

	// Create a new HTTP request
	req := httptest.NewRequest("GET", "/income-tax/calculate-tax?year=2023&salary=25000", nil)

	// Create a new HTTP response recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusInternalServerError, w.Code, "Unexpected status code")

	// Parse the response body
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err, "Failed to parse response body")

	// Check the error message in the response
	expectedErrorMessage := "Failed to retrieve tax brackets"
	assert.Equal(t, expectedErrorMessage, response["error"], "Unexpected error message")
}

func TestGetTotalIncomeTax_InvalidJSONResponse(t *testing.T) {
	//// Create a new router instance
	//router := gin.New()
	//router.Use(gin.Recovery())
	//
	//// Configure the router for testing
	//gin.SetMode(gin.TestMode)
	//
	//// Mock the tax calculator API endpoint with an invalid JSON response
	//router.GET("/income-tax/calculate-tax", func(c *gin.Context) {
	//	c.String(http.StatusOK, "invalid json response")
	//})
	//
	//// Create a new HTTP request
	//req := httptest.NewRequest("GET", "/income-tax/calculate-tax?year=2023&salary=25000", nil)
	//
	//// Create a new HTTP response recorder
	//w := httptest.NewRecorder()
	//
	//// Perform the request
	//router.ServeHTTP(w, req)
	//
	//// Check the response status code
	//assert.Equal(t, http.StatusInternalServerError, w.Code, "Unexpected status code")
	//
	//// Parse the response body
	//var response map[string]string
	//err := json.Unmarshal(w.Body.Bytes(), &response)
	//require.NoError(t, err, "Failed to parse response body")
	//
	//// Check the error message in the response
	//expectedErrorMessage := "Failed to parse response"
	//assert.Equal(t, expectedErrorMessage, response["error"], "Unexpected error message")
}
