package main

import (
	"github.com/joho/godotenv"
	"github.com/siparisa/interview-test-server/internal"
	"log"
	"os"
)

func main() {
	// Load the environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}

	logger := log.New(os.Stdout, "[main] ", log.LstdFlags)

	// Set up the router for the application
	router, err := internal.SetupRouter(logger)
	if err != nil {
		logger.Fatalf("Failed to set up router: %v", err)
	}

	// Get the APP port from the environment variable or use the default
	appPort := os.Getenv("PORT_APP")
	if appPort == "" {
		appPort = "8080"
	}

	// Run the application on the specified port
	addr := ":" + appPort
	logger.Printf("Starting server on port %s...", appPort)
	if err := router.Run(addr); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}
