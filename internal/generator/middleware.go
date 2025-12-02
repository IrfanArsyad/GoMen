package generator

import (
	"fmt"
	"path/filepath"
)

// MakeMiddleware generates a new middleware file
func MakeMiddleware(name string) {
	pascalName := toPascalCase(name)
	snakeName := toSnakeCase(pascalName)

	content := fmt.Sprintf(`package middlewares

import (
	"github.com/gin-gonic/gin"
)

func %sMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Before request
		// Add your middleware logic here

		// Process request
		c.Next()

		// After request
		// Add your post-processing logic here
	}
}
`, pascalName)

	filePath := filepath.Join(getProjectRoot(), "app", "middlewares", snakeName+".go")

	if err := writeFile(filePath, content); err != nil {
		printError(err)
		return
	}

	printSuccess("Middleware", filePath)
	fmt.Println("  → Don't forget to register this middleware in routes/api.go")
}
