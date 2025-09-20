package middleware

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/gin-gonic/gin"
)

func genReqID() string {
	b := make([]byte, 12)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetHeader("X-Request-Id")
		if id == "" {
			id = genReqID()
		}
		c.Request.Header.Set("X-Request-Id", id)
		c.Writer.Header().Set("X-Request-Id", id)
		c.Set("request_id", id)
		c.Next()
	}
}
