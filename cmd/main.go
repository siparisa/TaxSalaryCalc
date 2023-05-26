package main

import (
	"github.com/siparisa/interview-test-server/internal"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "[main] ", log.LstdFlags)

	// Set up the router for the application
	router, err := internal.SetupRouter(logger)
	if err != nil {
		logger.Fatalf("Failed to set up router: %v", err)
	}

	// Get the APP port from environment variable or use default
	appPort := os.Getenv("PORT_APP")
	if appPort == "" {
		appPort = "8080" // Use default port if environment variable is not set
	}

	// Run the application on the specified port
	addr := ":" + appPort
	logger.Printf("Starting server on port %s...", appPort)
	if err := router.Run(addr); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}
