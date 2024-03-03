build:
	@echo "Building..."
	@go build -o bin/app main.go

run: build
	@echo "Running..."
	@./bin/app