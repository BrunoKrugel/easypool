.PHONY: build
build: # Build the go application
	go build cmd/main.go

.PHONY: run
run: # Run the go application
	go run cmd/main.go

.PHONY: clean
clean: # Clean the go application
	go clean

.PHONY: all
all: 
	go build cmd/main.go
	go run cmd/main.go

.PHONY: test
test: # Run the go tests
	go test ./...
