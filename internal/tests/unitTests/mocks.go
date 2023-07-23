// tests/mocks.go

package tests

import (
	"errors"
	"github.com/siparisa/interview-test-server/internal/entity"
	"time"
)

// Define a mock tax service that implements the ITaxService interface.
type mockTaxService struct{}

func (m *mockTaxService) CalculateTaxForSalary(taxBrackets *entity.TaxBrackets, salary float64, totalTaxAmount float64) (float64, error) {
	if salary == 1 {
		return 0.15, nil
	} else if salary == 6000000000 {
		return 329999999979296.25, nil
	} else {
		return 7630.35, nil
	}
}

func (m *mockTaxService) CalculateTaxPerBand(taxBrackets *entity.TaxBrackets, salary float64) (*entity.TaxCalculationResult, error) {

	taxAmountPerBand := make(map[string]float64)
	taxAmountPerBand["band1"] = 0.15 * salary            // Tax rate for band1 is 0.15
	taxAmountPerBand["band2"] = 0.205 * (salary - 47630) // Tax rate for band2 is 0.205
	totalTaxAmount := 7630.35

	result := &entity.TaxCalculationResult{
		TaxAmountPerBand: taxAmountPerBand,
		TotalTaxAmount:   totalTaxAmount,
	}

	return result, nil
}

func (m *mockTaxService) CalculateEffectiveRate(taxAmount, salary float64) (float64, error) {
	return 15.26, nil
}

// Define a mock tax bracket service that implements the ITaxBracketService interface.
type mockTaxBracketService struct{}

func (m *mockTaxBracketService) GetTaxBracket(taxYear string, maxRetries int, retryInterval time.Duration) (*entity.TaxBrackets, error) {
	// Mock the tax brackets for the year 2019.
	if taxYear == "2019" {
		taxBrackets := &entity.TaxBrackets{
			TaxBrackets: []entity.TaxBracket{
				{
					Band: "band1",
					Max:  47630,
					Min:  0,
					Rate: 0.15,
				},
				{
					Band: "band2",
					Max:  95259,
					Min:  47630,
					Rate: 0.205,
				},
				{
					Band: "band3",
					Max:  147667,
					Min:  95259,
					Rate: 0.26,
				},
				{
					Band: "band4",
					Max:  210371,
					Min:  147667,
					Rate: 0.29,
				},
				{
					Band: "band5",
					Min:  210371,
					Rate: 0.33,
				},
			},
		}

		return taxBrackets, nil
	}

	// Return an error for other tax years.
	return nil, errors.New("tax brackets not found for the given year")
}
