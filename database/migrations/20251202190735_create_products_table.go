package migrations

import (
	"gorm.io/gorm"
)

// CreateProductsTable migration
type CreateProductsTable struct{}

func (m *CreateProductsTable) Up(db *gorm.DB) error {
	// Example: Create table
	// return db.Exec(`
	// 	CREATE TABLE IF NOT EXISTS table_name (
	// 		id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	// 		name VARCHAR(255) NOT NULL,
	// 		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	// 		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	// 		deleted_at TIMESTAMP NULL
	// 	)
	// `).Error

	// Or use GORM AutoMigrate
	// return db.AutoMigrate(&models.YourModel{})

	return nil
}

func (m *CreateProductsTable) Down(db *gorm.DB) error {
	// Example: Drop table
	// return db.Exec("DROP TABLE IF EXISTS table_name").Error

	return nil
}
