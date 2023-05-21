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
