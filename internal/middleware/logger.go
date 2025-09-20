package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		status := c.Writer.Status()
		reqID, _ := c.Get("request_id")
		clientIP := c.ClientIP()

		log.Printf("[req=%v] %s %s %d %s %s",
			reqID, c.Request.Method, c.Request.URL.Path, status, latency.String(), clientIP)
	}
}
