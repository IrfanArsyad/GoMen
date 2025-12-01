package main

import (
	"flag"
	"gomen/app/middlewares"
	"gomen/config"
	"gomen/database/migrations"
	"gomen/database/seeders"
	"gomen/routes"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Parse command line flags
	migrate := flag.Bool("migrate", false, "Run database migrations")
	seed := flag.Bool("seed", false, "Run database seeders")
	flag.Parse()

	// Load configuration
	config.Load()

	// Connect to database
	config.ConnectDatabase()

	// Run migrations if flag is set
	if *migrate {
		migrations.Migrate()
	}

	// Run seeders if flag is set
	if *seed {
		seeders.Seed()
	}

	// If only running migrations/seeds, exit
	if *migrate || *seed {
		if !*migrate && *seed {
			// If only seed flag, still need to continue to server
		} else if *migrate && !*seed {
			return
		} else if *migrate && *seed {
			return
		}
	}

	// Setup Gin
	if config.Get().App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// Global middlewares
	router.Use(middlewares.RecoveryMiddleware())
	router.Use(middlewares.LoggerMiddleware())
	router.Use(middlewares.CorsMiddleware())
	router.Use(middlewares.RateLimitMiddleware(100, time.Minute)) // 100 requests per minute

	// Setup routes
	routes.SetupRoutes(router)

	// Start server
	port := config.Get().App.Port
	log.Printf("Server starting on port %s", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server: " + err.Error())
	}
}
