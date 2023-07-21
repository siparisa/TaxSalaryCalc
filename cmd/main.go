package main

import (
	"log"
	"os"
)

func main() {
	router, logger, err := initializeApp()
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	appPort := os.Getenv("PORT_APP")
	if appPort == "" {
		appPort = "8080"
	}

	addr := ":" + appPort
	logger.Printf("Starting server on port %s...", appPort)
	if err := router.Run(addr); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}
