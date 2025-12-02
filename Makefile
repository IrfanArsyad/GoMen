.PHONY: build run dev migrate seed help clean controller model migration service request middleware seeder resource version list

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GORUN=$(GOCMD) run

# Build directory
BUILD_DIR=bin
BINARY_NAME=gomen

# Build the CLI tool
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/gomen
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

# Run the main application
run:
	@$(GORUN) main.go

# Run with hot reload (requires air: go install github.com/air-verse/air@latest)
dev:
	@air

# Run database migrations
migrate:
	@$(GORUN) main.go -migrate

# Run database seeders
seed:
	@$(GORUN) main.go -seed

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@$(GOCLEAN)
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete"

# ==================== Code Generators ====================

# Create a new controller: make controller name=Product
controller:
	@./bin/gomen make:controller $(name)

# Create a new model: make model name=Product
model:
	@./bin/gomen make:model $(name)

# Create a new migration: make migration name=create_products_table
migration:
	@./bin/gomen make:migration $(name)

# Create a new service: make service name=Product
service:
	@./bin/gomen make:service $(name)

# Create a new request: make request name=Product
request:
	@./bin/gomen make:request $(name)

# Create a new middleware: make middleware name=RateLimit
middleware:
	@./bin/gomen make:middleware $(name)

# Create a new seeder: make seeder name=Product
seeder:
	@./bin/gomen make:seeder $(name)

# Create a full resource (model, controller, service, request): make resource name=Product
resource:
	@./bin/gomen make:resource $(name)

# Show CLI version
version:
	@echo "GoMen CLI v1.0.0"

# Show all available commands (alias for help)
list:
	@echo "Available Commands:"
	@echo "  make:controller    Create a new controller"
	@echo "  make:model         Create a new model"
	@echo "  make:migration     Create a new migration file"
	@echo "  make:service       Create a new service"
	@echo "  make:request       Create a new request validation"
	@echo "  make:middleware    Create a new middleware"
	@echo "  make:seeder        Create a new seeder"
	@echo "  make:resource      Create model, controller, service, and request (full resource)"

# Show help
help:
	@echo "GoMen - Go REST API Starter Kit"
	@echo ""
	@echo "Usage:"
	@echo "  make build              Build the CLI tool"
	@echo "  make run                Run the application"
	@echo "  make dev                Run with hot reload (requires air)"
	@echo "  make migrate            Run database migrations"
	@echo "  make seed               Run database seeders"
	@echo "  make clean              Clean build artifacts"
	@echo ""
	@echo "Code Generators:"
	@echo "  make controller name=<Name>   Create a new controller"
	@echo "  make model name=<Name>        Create a new model"
	@echo "  make migration name=<name>    Create a new migration"
	@echo "  make service name=<Name>      Create a new service"
	@echo "  make request name=<Name>      Create a new request"
	@echo "  make middleware name=<Name>   Create a new middleware"
	@echo "  make seeder name=<Name>       Create a new seeder"
	@echo "  make resource name=<Name>     Create model, controller, service, request"
	@echo ""
	@echo "Examples:"
	@echo "  make controller name=Product"
	@echo "  make model name=Product"
	@echo "  make migration name=create_products_table"
	@echo "  make resource name=Product"
