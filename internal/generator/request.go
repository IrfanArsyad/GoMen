package generator

import (
	"fmt"
	"path/filepath"
)

// MakeRequest generates a new request validation file
func MakeRequest(name string) {
	pascalName := toPascalCase(name)
	snakeName := toSnakeCase(pascalName)

	content := fmt.Sprintf(`package requests

type Create%sRequest struct {
	Name string `+"`json:\"name\" validate:\"required,min=2,max=100\"`"+`
	// Add your validation fields here
}

type Update%sRequest struct {
	Name string `+"`json:\"name\" validate:\"required,min=2,max=100\"`"+`
	// Add your validation fields here
}
`, pascalName, pascalName)

	filePath := filepath.Join(getProjectRoot(), "app", "requests", snakeName+"_request.go")

	if err := writeFile(filePath, content); err != nil {
		printError(err)
		return
	}

	printSuccess("Request", filePath)
}
