package service

import (
	"errors"
	"github.com/shopspring/decimal"
	"github.com/siparisa/interview-test-server/internal/entity"
)

// ITaxService defines the interface for tax-related calculations.
type ITaxService interface {
	CalculateTaxForSalary(taxBrackets *entity.TaxBrackets, salary float64, totalTaxAmount float64) (float64, error)
	CalculateTaxPerBand(taxBrackets *entity.TaxBrackets, salary float64) (*entity.TaxCalculationResult, error)
	CalculateEffectiveRate(taxAmount, salary float64) (float64, error)
}

type taxService struct{}

// NewTaxService creates a new instance of the taxService.
func NewTaxService() ITaxService {
	return &taxService{}
}

func (s *taxService) CalculateTaxForSalary(taxBrackets *entity.TaxBrackets, salary float64, totalTaxAmount float64) (float64, error) {
	// Calculate the tax amount for the given salary based on the tax brackets.
	taxBracket, err := getTaxBracketForSalary(*taxBrackets, salary)
	if err != nil {
		return 0.0, err
	}

	// Use the totalTaxAmount if it's greater than 0, otherwise calculate the tax amount as before.
	var taxAmount float64
	if totalTaxAmount > 0 {
		taxAmount = totalTaxAmount
	} else {
		taxAmount = salary * taxBracket.Rate
	}

	return taxAmount, nil
}

// CalculateTaxPerBand calculates the tax amount per tax band based on the given salary and tax brackets.
func (s *taxService) CalculateTaxPerBand(taxBrackets *entity.TaxBrackets, salary float64) (*entity.TaxCalculationResult, error) {
	taxAmountPerBand := make(map[string]float64)
	totalTaxAmount := 0.0

	for i, bracket := range taxBrackets.TaxBrackets {
		taxableIncome := 0.0

		if i == len(taxBrackets.TaxBrackets)-1 {
			// Handle the last tax bracket separately
			taxableIncome = salary - bracket.Min
		} else if salary > bracket.Max {
			taxableIncome = bracket.Max - bracket.Min
		} else if salary > bracket.Min {
			taxableIncome = salary - bracket.Min
		}

		if taxableIncome > 0 {
			taxAmount := decimal.NewFromFloat(taxableIncome).Mul(decimal.NewFromFloat(bracket.Rate))
			roundedAmount, _ := taxAmount.Round(2).Float64()
			taxAmountPerBand[bracket.Band] = roundedAmount
			totalTaxAmount += roundedAmount
		}
	}

	// Check for negative total tax amount
	if totalTaxAmount < 0 {
		return nil, errors.New("total tax amount cannot be negative")
	}

	// Create and return the custom entity
	result := &entity.TaxCalculationResult{
		TaxAmountPerBand: taxAmountPerBand,
		TotalTaxAmount:   totalTaxAmount,
	}

	return result, nil
}

// CalculateEffectiveRate calculates the effective tax rate based on the given tax amount and salary.
func (s *taxService) CalculateEffectiveRate(taxAmount, salary float64) (float64, error) {
	if salary == 0.0 {
		return 0.0, errors.New("salary cannot be zero")
	}

	taxRate := decimal.NewFromFloat(taxAmount).Div(decimal.NewFromFloat(salary)).Mul(decimal.NewFromFloat(100))
	roundedRate, _ := taxRate.Round(2).Float64() // Round to 2 decimal places and convert to float64
	return roundedRate, nil
}

// getTaxBracketForSalary finds the appropriate tax bracket for the given salary.
func getTaxBracketForSalary(taxBrackets entity.TaxBrackets, salary float64) (*entity.TaxBracket, error) {
	for _, bracket := range taxBrackets.TaxBrackets {
		if salary >= bracket.Min && salary <= bracket.Max {
			return &bracket, nil
		}
	}
	return nil, errors.New("tax bracket not found for the given salary")
}
