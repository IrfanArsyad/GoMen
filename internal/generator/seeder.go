package generator

import (
	"fmt"
	"path/filepath"
)

// MakeSeeder generates a new seeder file
func MakeSeeder(name string) {
	pascalName := toPascalCase(name)
	snakeName := toSnakeCase(pascalName)
	camelName := toCamelCase(pascalName)

	content := fmt.Sprintf(`package seeders

import (
	"gomen/app/models"
	"gomen/config"
	"log"
)

func Seed%s() {
	db := config.GetDB()

	%s := []models.%s{
		{
			Name: "Sample %s 1",
			// Add more fields here
		},
		{
			Name: "Sample %s 2",
			// Add more fields here
		},
	}

	for _, item := range %s {
		if err := db.FirstOrCreate(&item, models.%s{Name: item.Name}).Error; err != nil {
			log.Printf("Failed to seed %s: %%v", err)
		}
	}

	log.Println("%s seeded successfully")
}
`, pascalName, camelName, pascalName, pascalName, pascalName, camelName, pascalName, snakeName, pascalName)

	filePath := filepath.Join(getProjectRoot(), "database", "seeders", snakeName+"_seeder.go")

	if err := writeFile(filePath, content); err != nil {
		printError(err)
		return
	}

	printSuccess("Seeder", filePath)
	fmt.Println("  → Don't forget to call this seeder in database/seeders/seeder.go")
}
