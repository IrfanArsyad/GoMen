package generator

import (
	"fmt"
)

// MakeResource generates a complete resource (model, controller, service, request)
func MakeResource(name string) {
	pascalName := toPascalCase(name)

	fmt.Printf("\n🚀 Creating resource: %s\n\n", pascalName)

	// Create Model
	MakeModel(name)

	// Create Request
	MakeRequest(name)

	// Create Service
	MakeService(name)

	// Create Controller
	MakeController(name)

	fmt.Println("\n✨ Resource created successfully!")
	fmt.Println("\nNext steps:")
	fmt.Printf("  1. Update the model fields in app/models/%s.go\n", toSnakeCase(pascalName))
	fmt.Printf("  2. Update the request validation in app/requests/%s_request.go\n", toSnakeCase(pascalName))
	fmt.Printf("  3. Update the service logic in app/services/%s_service.go\n", toSnakeCase(pascalName))
	fmt.Printf("  4. Add the model to database/migrations/migrate.go\n")
	fmt.Printf("  5. Register routes in routes/api.go\n")
}
