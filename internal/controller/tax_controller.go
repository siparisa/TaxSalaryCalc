package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/siparisa/interview-test-server/internal/controller/helper"
	_ "github.com/siparisa/interview-test-server/internal/entity"
	"github.com/siparisa/interview-test-server/internal/service"
	"strconv"
)

type TaxController struct {
	taxService        service.ITaxService
	taxBracketService service.ITaxBracketService
}

// NewTaxController creates a new instance of TaxController with the given ITaxService and ITaxBracketService.
// It takes ITaxService and ITaxBracketService as parameters and returns a pointer to the newly created TaxController.
func NewTaxController(taxService service.ITaxService, taxBracketService service.ITaxBracketService) *TaxController {
	return &TaxController{
		taxService:        taxService,
		taxBracketService: taxBracketService,
	}
}

// GetTotalIncomeTax @Summary Get total income tax
// @Description Calculate the total income tax based on salary and tax year
// @ID getTotalIncomeTax
// @Accept json
// @Produce json
// @Param salary query string true "Salary"
// @Param year query string true "Tax Year"
// @Success 200 {object} TaxAmountResponse
// @Failure 400 {object} APIError
// @Router /calculate-tax [get]
func (c *TaxController) GetTotalIncomeTax(ctx *gin.Context) {
	var qp helper.GetIncomeTaxParams
	if err := ctx.ShouldBindQuery(&qp); err != nil {

		if ve, ok := err.(validator.ValidationErrors); ok {
			errorMsg := ""
			for _, e := range ve {
				errorMsg += fmt.Sprintf("Field %s failed validation: %s\n", e.Field(), e.Tag())
			}
			helper.BadRequest(ctx, errorMsg)
			return
		}

		helper.BadRequest(ctx, "Invalid query parameters")
		return
	}

	salaryStr := ctx.Query("salary")
	taxYear := ctx.Query("year")

	salary, err := strconv.ParseFloat(salaryStr, 64)
	if err != nil {
		helper.InternalServerError(ctx, "Invalid salary")
		return
	}

	// Retrieve the tax brackets for the given year
	taxBrackets, err := c.taxBracketService.GetTaxBracket(taxYear)
	if err != nil {
		helper.InternalServerError(ctx, "Failed to get tax brackets")
		return
	}

	// Calculate the tax amount for the salary
	taxAmount, err := c.taxService.CalculateTaxForSalary(taxBrackets, salary)
	if err != nil {
		helper.InternalServerError(ctx, "Failed to calculate tax amount")
		return
	}

	// Calculate the tax amount per band
	taxAmountPerBand, err := c.taxService.CalculateTaxPerBand(taxBrackets, salary)
	if err != nil {
		helper.InternalServerError(ctx, "Failed to calculate tax amount per band")
		return
	}

	// Calculate the effective tax rate
	effectiveRate, err := c.taxService.CalculateEffectiveRate(taxAmount, salary)
	if err != nil {
		helper.InternalServerError(ctx, "Failed to calculate Effective Rate")
		return
	}

	// Prepare the response
	response := helper.TaxAmountResponse{
		TotalTaxAmount:   taxAmount,
		TaxAmountPerBand: taxAmountPerBand,
		EffectiveRate:    effectiveRate,
	}

	helper.OK(ctx, response)
}
