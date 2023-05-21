package main

import (
	"fmt"
	"github.com/siparisa/interview-test-server/internal/config"
	"log"
	"os"
)

func main() {

	logger := log.New(os.Stdout, "[main] ", log.LstdFlags)

	// Set up the router for the application
	router, err := setupRouter(logger)
	if err != nil {
		logger.Fatalf("Failed to set up router: %v", err)
	}

	// Get the APP port from environment variable or use default
	appPort := config.GetPort("PORT_APP", config.AppDefaultPort)

	// Run the application on the specified port
	addr := fmt.Sprintf(":%s", appPort)
	logger.Printf("Starting server on port %s...", appPort)
	if err := router.Run(addr); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}
