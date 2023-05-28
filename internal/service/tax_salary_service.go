package service

import (
	"github.com/siparisa/interview-test-server/internal/entity"
)

// getTaxBracketForSalary finds the appropriate tax bracket for the given salary.
func getTaxBracketForSalary(taxBrackets entity.TaxBrackets, salary float64) *entity.TaxBracket {
	for _, bracket := range taxBrackets.TaxBrackets {
		if salary >= bracket.Min && salary <= bracket.Max {
			return &bracket
		}
	}
	return nil
}

// CalculateTaxForSalary calculates the tax amount for the given salary based on the tax brackets.
func CalculateTaxForSalary(taxResponse *entity.TaxBrackets, salary float64) float64 {
	taxBracket := getTaxBracketForSalary(*taxResponse, salary)
	if taxBracket != nil {
		taxAmount := salary * taxBracket.Rate
		return taxAmount
	}
	return 0.0
}

// CalculateTaxPerBand calculates the tax amount per tax band based on the given salary and tax brackets.
func CalculateTaxPerBand(taxBrackets *entity.TaxBrackets, salary float64) map[string]float64 {
	taxAmountPerBand := make(map[string]float64)
	previousMax := 0.0

	for _, bracket := range taxBrackets.TaxBrackets {
		if salary > bracket.Max {
			taxableIncome := bracket.Max - previousMax
			taxAmount := taxableIncome * bracket.Rate
			taxAmountPerBand[bracket.Band] = taxAmount
		} else if salary > bracket.Min {
			taxableIncome := salary - previousMax
			taxAmount := taxableIncome * bracket.Rate
			taxAmountPerBand[bracket.Band] = taxAmount
		}

		previousMax = bracket.Max
	}

	return taxAmountPerBand
}

// CalculateEffectiveRate calculates the effective tax rate based on the given tax amount and salary.
func CalculateEffectiveRate(taxAmount, salary float64) float64 {
	return (taxAmount / salary) * 100
}
