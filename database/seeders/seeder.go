package seeders

import (
	"gomen/app/models"
	"gomen/config"
	"gomen/helpers"
	"log"
)

func Seed() {
	log.Println("Running database seeders...")

	// Seed admin user
	seedAdminUser()

	log.Println("Database seeding completed successfully")
}

func seedAdminUser() {
	db := config.GetDB()
	hashedPassword, _ := helpers.HashPassword("password123")

	admin := models.User{
		Name:     "Admin",
		Email:    "admin@example.com",
		Password: hashedPassword,
		IsActive: true,
	}

	if err := db.FirstOrCreate(&admin, models.User{Email: admin.Email}).Error; err != nil {
		log.Printf("Failed to seed admin user: %v", err)
	} else {
		log.Println("Admin user seeded successfully")
	}
}
