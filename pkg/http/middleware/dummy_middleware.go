package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

// DummyMiddleware is an example Gin middleware
func DummyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Before request
		start := time.Now()
		slog.Info("DummyMiddleware: request started")

		// Call the next handler
		c.Next()

		// After request
		duration := time.Since(start)
		status := c.Writer.Status()

		slog.Info("DummyMiddleware: request finished",
			"status", status,
			"duration", duration,
		)
	}
}
