run:
    # Run the Application
	go run ./cmd/*.go

test:
	# Run the tests
	go test ./internal/tests/...
