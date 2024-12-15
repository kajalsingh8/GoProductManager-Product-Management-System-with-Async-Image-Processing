package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		duration := time.Since(startTime)
		log.Printf("Request: %s %s | Duration: %s", c.Request.Method, c.Request.URL.Path, duration)
	}
}
