package generator

import (
	"fmt"
	"path/filepath"
)

// MakeModel generates a new model file
func MakeModel(name string) {
	pascalName := toPascalCase(name)
	snakeName := toSnakeCase(pascalName)
	tableName := toPlural(snakeName)

	content := fmt.Sprintf(`package models

type %s struct {
	BaseModel
	Name string `+"`json:\"name\" gorm:\"size:255;not null\"`"+`
	// Add your fields here
}

func (%s) TableName() string {
	return "%s"
}
`, pascalName, pascalName, tableName)

	filePath := filepath.Join(getProjectRoot(), "app", "models", snakeName+".go")

	if err := writeFile(filePath, content); err != nil {
		printError(err)
		return
	}

	printSuccess("Model", filePath)
	fmt.Println("  → Don't forget to add this model to database/migrations/migrate.go")
}
