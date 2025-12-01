package middlewares

import (
	"gomen/app/responses"
	"gomen/helpers"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			responses.Unauthorized(c, "Authorization header is required")
			c.Abort()
			return
		}

		// Check Bearer token format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			responses.Unauthorized(c, "Invalid authorization header format")
			c.Abort()
			return
		}

		token := parts[1]

		// Validate JWT token
		claims, err := helpers.ValidateJWT(token)
		if err != nil {
			responses.Unauthorized(c, "Invalid or expired token")
			c.Abort()
			return
		}

		// Set user info to context
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)

		c.Next()
	}
}

// OptionalAuthMiddleware - middleware that allows unauthenticated requests
func OptionalAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.Next()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.Next()
			return
		}

		token := parts[1]
		claims, err := helpers.ValidateJWT(token)
		if err == nil {
			c.Set("user_id", claims.UserID)
			c.Set("email", claims.Email)
		}

		c.Next()
	}
}
