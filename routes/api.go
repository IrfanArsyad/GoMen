package routes

import (
	"gomen/app/controllers"
	"gomen/app/middlewares"
	"gomen/app/responses"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Health check
	router.GET("/health", func(c *gin.Context) {
		responses.Success(c, "OK", gin.H{
			"status": "healthy",
		})
	})

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		setupAuthRoutes(v1)
		setupUserRoutes(v1)
	}
}

func setupAuthRoutes(rg *gin.RouterGroup) {
	authController := controllers.NewAuthController()

	auth := rg.Group("/auth")
	{
		// Public routes
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)

		// Protected routes
		protected := auth.Group("")
		protected.Use(middlewares.AuthMiddleware())
		{
			protected.GET("/profile", authController.GetProfile)
			protected.PUT("/profile", authController.UpdateProfile)
			protected.POST("/change-password", authController.ChangePassword)
			protected.POST("/refresh", authController.RefreshToken)
		}
	}
}

func setupUserRoutes(rg *gin.RouterGroup) {
	userController := controllers.NewUserController()

	users := rg.Group("/users")
	users.Use(middlewares.AuthMiddleware())
	{
		users.GET("", userController.Index)
		users.GET("/:id", userController.Show)
		users.POST("", userController.Store)
		users.PUT("/:id", userController.Update)
		users.DELETE("/:id", userController.Delete)
	}
}
