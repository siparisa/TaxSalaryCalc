package main

import (
	"log"
	"os"

	"github.com/siparisa/interview-test-server/internal/controller"
	"github.com/siparisa/interview-test-server/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/siparisa/interview-test-server/internal"
)

// initializeApp sets up the application by initializing services, controllers, and the router.
// It reads environment variables, such as TAX_CALCULATOR_URL, to configure the application.
// It returns the Gin router, a logger instance, and any error encountered during initialization.
func initializeApp() (*gin.Engine, *log.Logger, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, nil, err
	}

	logger := log.New(os.Stdout, "[main] ", log.LstdFlags)

	taxCalculatorURL := os.Getenv("TAX_CALCULATOR_URL")
	if taxCalculatorURL == "" {
		taxCalculatorURL = "http://localhost:7070/tax-calculator/tax-year/" // Default URL if not provided
	}

	taxService := service.NewTaxService()
	taxBracketService := service.NewTaxBracketService(taxCalculatorURL)
	taxController := controller.NewTaxController(taxService, taxBracketService)

	router, err := internal.SetupRouter(logger, taxController)
	if err != nil {
		return nil, nil, err
	}

	return router, logger, nil
}
