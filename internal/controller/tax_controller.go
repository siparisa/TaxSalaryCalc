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

	taxBracket, err := service.GetTaxBracket(taxYear)
	if err != nil {
		helper.InternalServerError(ctx, "Failed to get tax bracket", "Failed to get tax bracket.")
		return
	}

	salary, err := strconv.ParseFloat(salaryStr, 64)
	if err != nil {
		helper.InternalServerError(ctx, "Invalid salary", "salary value is not numeric.")
		return
	}

	taxAmount := service.CalculateTaxForSalary(taxBracket, salary)

	helper.OK(ctx, gin.H{
		"taxAmount": taxAmount,
	})
}
