package migrations

import (
	"gomen/app/models"
	"gomen/config"
	"log"
)

func Migrate() {
	db := config.GetDB()

	log.Println("Running database migrations...")

	err := db.AutoMigrate(
		&models.User{},
		&models.Product{},
	)

	if err != nil {
		log.Fatal("Migration failed: " + err.Error())
	}

	log.Println("Database migrations completed successfully")
}
