package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/siparisa/interview-test-server/internal/controller"
	"github.com/siparisa/interview-test-server/internal/controller/helper"
)

func TestGetTotalIncomeTax(t *testing.T) {
	router := gin.New()
	taxService := &mockTaxService{}
	taxBracketService := &mockTaxBracketService{}
	taxController := controller.NewTaxController(taxService, taxBracketService)
	router.Handle(http.MethodGet, "/calculate-tax", taxController.GetTotalIncomeTax)

	t.Run("InvalidSalaryInput", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/calculate-tax?salary=invalid&year=2019", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)
		if rec.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("NegativeSalaryInput", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/calculate-tax?salary=-50000&year=2019", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)
		if rec.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("InvalidTaxYearInput", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/calculate-tax?salary=50000&year=invalid", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)
		if rec.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Success", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/calculate-tax?salary=50000&year=2019", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, rec.Code)
		}

		var response helper.TaxAmountResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal response: %v", err)
		}

		expectedTotalTaxAmount := 7630.35
		if response.TotalTaxAmount != expectedTotalTaxAmount {
			t.Errorf("Expected total tax amount %f, but got %f", expectedTotalTaxAmount, response.TotalTaxAmount)
		}
	})

	t.Run("TestGetTotalIncomeTaxWithMissingQueryParameters", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/calculate-tax", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)
		if rec.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("TestGetTotalIncomeTaxWithLargeSalary", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/calculate-tax?salary=1000000000000000&year=2019", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, rec.Code)
		}
	})

	t.Run("TestGetTotalIncomeTaxWithSmallSalary", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/calculate-tax?salary=1&year=2019", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, rec.Code)
		}

		var response helper.TaxAmountResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal response: %v", err)
		}

		expectedTotalTaxAmount := 0.15
		if response.TotalTaxAmount != expectedTotalTaxAmount {
			t.Errorf("Expected total tax amount %f, but got %f", expectedTotalTaxAmount, response.TotalTaxAmount)
		}
	})

	t.Run("TestGetTotalIncomeTaxWithNonNumericSalary", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/calculate-tax?salary=abc&year=2019", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)
		if rec.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, rec.Code)
		}
	})
}
