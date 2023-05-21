run:
	# Prompt the user to enter the port on which the Docker is running
	read -p "Enter the port which the Docker is running on: " PORT_TAX_YEAR; \

	# Prompt the user to enter the port for calling the API
	read -p "Enter the port that you want to call the API: " PORT_APP; \

	# Run the application with the specified ports
	go run -ldflags="-s -w" -race -tags=dev -ldflags "$(BUILD_INFO)" ./internal/main.go ./internal/server.go -port1 $$PORT_TAX_YEAR -port2 $$PORT_APP

test:
	# Run the tests
	go test ./internal/tests/...
