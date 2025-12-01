package middlewares

import (
	"gomen/app/responses"
	"log"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v\n%s", err, debug.Stack())
				responses.InternalServerError(c, "Internal server error")
				c.Abort()
			}
		}()
		c.Next()
	}
}
