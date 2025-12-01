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
		// Add more models here
	)

	if err != nil {
		log.Fatal("Migration failed: " + err.Error())
	}

	log.Println("Database migrations completed successfully")
}
