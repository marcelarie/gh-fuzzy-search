# Run the Go app
run ARGS="":
    go run . {{ARGS}}

# Build the Go app
build ARGS="":
    go build {{ARGS}}

# Format the Go app
format:
    go fmt

# Test the Go app
test:
    go test
