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

// APIError represents the JSON response for API errors
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// BadRequest sends a Bad Request response with the provided error message.
func BadRequest(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusBadRequest, APIError{
		Code:    http.StatusBadRequest,
		Message: message,
	})
}

// InternalServerError sends an Internal Server Error response with the provided error message.
func InternalServerError(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusInternalServerError, APIError{
		Code:    http.StatusInternalServerError,
		Message: message,
	})
}

// OK sends a success response with the provided data.
func OK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}
