package service

import (
	"errors"
	"github.com/shopspring/decimal"
	"github.com/siparisa/interview-test-server/internal/entity"
)

// ITaxService defines the interface for tax-related calculations.
type ITaxService interface {
	CalculateTaxForSalary(taxResponse *entity.TaxBrackets, salary float64) (float64, error)
	CalculateTaxPerBand(taxBrackets *entity.TaxBrackets, salary float64) (map[string]float64, error)
	CalculateEffectiveRate(taxAmount, salary float64) (float64, error)
}

type taxService struct{}

// NewTaxService creates a new instance of the taxService.
func NewTaxService() ITaxService {
	return &taxService{}
}

// CalculateTaxForSalary calculates the tax amount for the given salary based on the tax brackets.
func (s *taxService) CalculateTaxForSalary(taxResponse *entity.TaxBrackets, salary float64) (float64, error) {
	taxBracket := getTaxBracketForSalary(*taxResponse, salary)
	if taxBracket != nil {
		taxAmount := decimal.NewFromFloat(salary).Mul(decimal.NewFromFloat(taxBracket.Rate))
		roundedAmount, _ := taxAmount.Round(2).Float64() // Round to 2 decimal places and convert to float64
		return roundedAmount, nil
	}
	return 0.0, errors.New("tax bracket not found for the given salary")
}

// CalculateTaxPerBand calculates the tax amount per tax band based on the given salary and tax brackets.
func (s *taxService) CalculateTaxPerBand(taxBrackets *entity.TaxBrackets, salary float64) (map[string]float64, error) {
	taxAmountPerBand := make(map[string]float64)
	previousMax := 0.0

	for _, bracket := range taxBrackets.TaxBrackets {
		if salary > bracket.Max {
			taxableIncome := bracket.Max - previousMax
			taxAmount := decimal.NewFromFloat(taxableIncome).Mul(decimal.NewFromFloat(bracket.Rate))
			roundedAmount, _ := taxAmount.Round(2).Float64() // Round to 2 decimal places and convert to float64
			taxAmountPerBand[bracket.Band] = roundedAmount
		} else if salary > bracket.Min {
			taxableIncome := salary - previousMax
			taxAmount := decimal.NewFromFloat(taxableIncome).Mul(decimal.NewFromFloat(bracket.Rate))
			roundedAmount, _ := taxAmount.Round(2).Float64() // Round to 2 decimal places and convert to float64
			taxAmountPerBand[bracket.Band] = roundedAmount
		}

		previousMax = bracket.Max
	}

	return taxAmountPerBand, nil
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
func getTaxBracketForSalary(taxBrackets entity.TaxBrackets, salary float64) *entity.TaxBracket {
	for _, bracket := range taxBrackets.TaxBrackets {
		if salary >= bracket.Min && salary <= bracket.Max {
			return &bracket
		}
	}
	return nil
}
