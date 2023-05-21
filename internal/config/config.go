package config

import (
	"os"
)

var (
	TaxYearDefaultPort = "7070" // Default port for tax year if not specified
	AppDefaultPort     = "8080" // Default port for API if not specified
)

func GetPort(portName string, defaultPort string) string {
	port := os.Getenv(portName)
	if port == "" {
		return defaultPort
	}
	return port
}
