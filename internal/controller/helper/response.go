package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// TaxAmountResponse represents the response for the calculate-tax endpoint
type TaxAmountResponse struct {
	TotalTaxAmount   float64            `json:"totalTaxAmount"`
	TaxAmountPerBand map[string]float64 `json:"taxAmountPerBand"`
	EffectiveRate    float64            `json:"effectiveRate"`
}

// BadRequest sends a Bad Request response with the provided error message.
func BadRequest(ctx *gin.Context, message string, details string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"details": details,
		"error":   message,
	})
}

// InternalServerError sends an Internal Server Error response with the provided error message.
func InternalServerError(ctx *gin.Context, message string, details string) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error":   message,
		"details": details,
	})
}

// BadRequestWithDetails sends a Bad Request response with the provided error message and details.
func BadRequestWithDetails(ctx *gin.Context, message, details string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error":   message,
		"details": details,
	})
}

// OK sends a success response with the provided data.
func OK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}
