package helper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetIncomeTaxParams is  query params for getting salary and year to calculate tax
type GetIncomeTaxParams struct {
	Salary string `form:"salary" binding:"required,numeric"`
	Year   string `form:"year" binding:"required,numeric,len=4"`
}

// IsValidTaxYear checks if the provided tax year is valid.
// It returns true if the tax year is valid, otherwise false.
func IsValidTaxYear(taxYear string) bool {
	validTaxYears := []string{"2019", "2020", "2021", "2022"}
	for _, year := range validTaxYears {
		if year == taxYear {
			return true
		}
	}
	return false
}

// IsValidSalary checks if the given salary string is valid (non-negative float64).
// If the salary is valid, it returns the parsed salary value. Otherwise, it returns an error.
func IsValidSalary(ctx *gin.Context, salaryStr string) (float64, error) {
	salary, err := strconv.ParseFloat(salaryStr, 64)
	if err != nil {
		InternalServerError(ctx, "Invalid salary")
		return 0, err
	}

	if salary < 0 {
		BadRequest(ctx, "Salary cannot be negative")
		return 0, fmt.Errorf("salary cannot be negative")
	}

	return salary, nil
}
