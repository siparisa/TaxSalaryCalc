package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/siparisa/interview-test-server/internal/controller/helper"
	"github.com/siparisa/interview-test-server/internal/service"
	"strconv"
)

// GetTotalIncomeTax @Summary Get total income tax
// @Description Calculate the total income tax based on salary and tax year
// @ID getTotalIncomeTax
// @Accept json
// @Produce json
// @Param salary query string true "Salary"
// @Param year query string true "Tax Year"
// @Success 200 {object} TaxAmountResponse
// @Failure 400 {object} ErrorResponse
// @Router /calculate-tax [get]
func GetTotalIncomeTax(ctx *gin.Context) {
	var qp helper.GetIncomeTaxParams
	if err := ctx.ShouldBindQuery(&qp); err != nil {
		helper.BadRequest(ctx, "Missing query parameters", "year and salary are mandatory params.")
		return
	}

	salaryStr := ctx.Query("salary")
	taxYear := ctx.Query("year")

	salary, err := strconv.ParseFloat(salaryStr, 64)
	if err != nil {
		helper.InternalServerError(ctx, "Invalid salary", "Salary value is not numeric.")
		return
	}

	// Retrieve the tax brackets for the given year
	taxBrackets, err := service.GetTaxBracket(taxYear)
	if err != nil {
		helper.InternalServerError(ctx, "Failed to get tax brackets", "Failed to retrieve tax brackets.")
		return
	}

	// Calculate the tax amount for the salary
	taxAmount := service.CalculateTaxForSalary(taxBrackets, salary)

	// Calculate the tax amount per band
	taxAmountPerBand := service.CalculateTaxPerBand(taxBrackets, salary)

	// Calculate the effective tax rate
	effectiveRate := service.CalculateEffectiveRate(taxAmount, salary)

	// Prepare the response
	response := helper.TaxAmountResponse{
		TotalTaxAmount:   taxAmount,
		TaxAmountPerBand: taxAmountPerBand,
		EffectiveRate:    effectiveRate,
	}

	helper.OK(ctx, response)
}
