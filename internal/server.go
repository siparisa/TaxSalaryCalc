package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/siparisa/interview-test-server/internal/controller"
	"log"
)

// SetupRouter creates and configures the Gin router for the application with the provided logger
func SetupRouter(logger *log.Logger, taxController *controller.TaxController) (*gin.Engine, error) {
	router := gin.Default()

	// Create a router group for "income-tax" endpoints
	incomeTaxGroup := router.Group("/income-tax")

	incomeTaxGroup.GET("/calculate-tax", func(c *gin.Context) {
		logger.Println("Handling GET request for /income-tax/calculate-tax")
		taxController.GetTotalIncomeTax(c)
	})

	return router, nil
}
