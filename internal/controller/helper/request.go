package helper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
)

// GetIncomeTaxParams is  query params for getting salary and year to calculate tax
type GetIncomeTaxParams struct {
	Salary string `form:"salary" binding:"required,numeric"`
	Year   string `form:"year" binding:"required,numeric,len=4"`
}

// GetValidationErrorMessage generates the validation error message for the provided validation errors.
func GetValidationErrorMessage(ve validator.ValidationErrors) string {
	var errorMsgSalary, errorMsgYear string
	for _, e := range ve {
		switch e.Field() {
		case "Salary":
			switch e.Tag() {
			case "required":
				errorMsgSalary = "Salary is required"
			case "numeric":
				errorMsgSalary = "Salary must be a numeric value"
			default:
				errorMsgSalary = "Invalid Salary"
			}
		case "Year":
			switch e.Tag() {
			case "required":
				errorMsgYear = "Year is required"
			case "numeric":
				errorMsgYear = "Year must be a numeric value"
			case "len":
				errorMsgYear = "Year must be exactly 4 digits long"
			default:
				errorMsgSalary = "Invalid Year"
			}
		default:
			errorMsgSalary = "Invalid Year and Salary Parameters"
		}
	}

	// Combine the error messages for Salary and Year
	errorMsg := errorMsgSalary + "\n" + errorMsgYear

	if errorMsg == "\n" {
		// If there are no specific error messages, use a generic one
		errorMsg = "Invalid query parameters"
	}

	return errorMsg
}

// IsValidTaxYear checks if the provided tax year is valid.
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
