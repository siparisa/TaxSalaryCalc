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
	var taxAmount decimal.Decimal
	if totalTaxAmount > 0 {
		taxAmount = decimal.NewFromFloat(totalTaxAmount)
	} else {
		taxAmount = decimal.NewFromFloat(salary).Mul(decimal.NewFromFloat(taxBracket.Rate))
	}

	// Convert the decimal.Decimal taxAmount back to float64 for compatibility with the existing code.
	roundedAmount, _ := taxAmount.Round(2).Float64()

	return roundedAmount, nil
}

// CalculateTaxPerBand calculates the tax amount per tax band based on the given salary and tax brackets.
func (s *taxService) CalculateTaxPerBand(taxBrackets *entity.TaxBrackets, salary float64) (*entity.TaxCalculationResult, error) {
	taxAmountPerBand := make(map[string]float64)
	totalTaxAmount := decimal.NewFromFloat(0)

	for i, bracket := range taxBrackets.TaxBrackets {
		taxableIncome := decimal.NewFromFloat(0)

		if i == len(taxBrackets.TaxBrackets)-1 {
			// Handle the last tax bracket separately
			taxableIncome = decimal.NewFromFloat(salary).Sub(decimal.NewFromFloat(bracket.Min))
		} else if salary > bracket.Max {
			taxableIncome = decimal.NewFromFloat(bracket.Max).Sub(decimal.NewFromFloat(bracket.Min))
		} else if salary > bracket.Min {
			taxableIncome = decimal.NewFromFloat(salary).Sub(decimal.NewFromFloat(bracket.Min))
		}

		if taxableIncome.GreaterThan(decimal.NewFromFloat(0)) {
			taxAmount := taxableIncome.Mul(decimal.NewFromFloat(bracket.Rate))
			roundedAmount, _ := taxAmount.Round(2).Float64()
			taxAmountPerBand[bracket.Band] = roundedAmount
			totalTaxAmount = totalTaxAmount.Add(decimal.NewFromFloat(roundedAmount))
		}
	}

	// Check for negative total tax amount
	if totalTaxAmount.LessThan(decimal.NewFromFloat(0)) {
		return nil, errors.New("total tax amount cannot be negative")
	}

	// Convert the decimal.Decimal totalTaxAmount back to float64 for compatibility with the existing code.
	roundedTotalAmount, _ := totalTaxAmount.Round(2).Float64()

	// Create and return the custom entity
	result := &entity.TaxCalculationResult{
		TaxAmountPerBand: taxAmountPerBand,
		TotalTaxAmount:   roundedTotalAmount,
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
	decimalSalary := decimal.NewFromFloat(salary)

	for _, bracket := range taxBrackets.TaxBrackets {
		decimalMin := decimal.NewFromFloat(bracket.Min)

		// Check if the current bracket is the last one (no maximum value specified)
		if bracket.Max == 0 {
			// The last bracket covers all salaries greater than or equal to its minimum value
			if decimalSalary.GreaterThanOrEqual(decimalMin) {
				return &bracket, nil
			}
		} else {
			decimalMax := decimal.NewFromFloat(bracket.Max)

			// Check if the salary falls within the current bracket's range
			if decimalSalary.GreaterThanOrEqual(decimalMin) && decimalSalary.LessThanOrEqual(decimalMax) {
				return &bracket, nil
			}
		}
	}

	return nil, errors.New("tax bracket not found for the given salary")
}
