package generator

import (
	"fmt"
	"path/filepath"
	"time"
)

// MakeMigration generates a new migration file
func MakeMigration(name string) {
	snakeName := toSnakeCase(name)
	timestamp := time.Now().Format("20060102150405")
	fileName := fmt.Sprintf("%s_%s", timestamp, snakeName)

	content := fmt.Sprintf(`package migrations

import (
	"gorm.io/gorm"
)

// %s migration
type %s struct{}

func (m *%s) Up(db *gorm.DB) error {
	// Example: Create table
	// return db.Exec(`+"`"+`
	// 	CREATE TABLE IF NOT EXISTS table_name (
	// 		id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	// 		name VARCHAR(255) NOT NULL,
	// 		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	// 		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	// 		deleted_at TIMESTAMP NULL
	// 	)
	// `+"`"+`).Error

	// Or use GORM AutoMigrate
	// return db.AutoMigrate(&models.YourModel{})

	return nil
}

func (m *%s) Down(db *gorm.DB) error {
	// Example: Drop table
	// return db.Exec("DROP TABLE IF EXISTS table_name").Error

	return nil
}
`, toPascalCase(snakeName), toPascalCase(snakeName), toPascalCase(snakeName), toPascalCase(snakeName))

	filePath := filepath.Join(getProjectRoot(), "database", "migrations", fileName+".go")

	if err := writeFile(filePath, content); err != nil {
		printError(err)
		return
	}

	printSuccess("Migration", filePath)
	fmt.Println("  → Register this migration in database/migrations/migrate.go")
}
