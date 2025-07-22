# Simple Golang Calculator with Testify

A simple Go calculator demo with Testify for testing.

## Installation

```bash
# Install dependencies
go mod tidy
```

## Running Tests

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...
```

## Generating Test Reports

### JUnit XML Report

```bash
# Install gotestsum (only needed once)
go install gotest.tools/gotestsum@latest

# Make sure ~/go/bin is in your PATH
export PATH=$PATH:~/go/bin

# Generate JUnit XML report
gotestsum --junitfile test-results.xml -- ./...
```

### JSON Report

```bash
# Generate JSON report
go test -json ./... > test-results.json
```
