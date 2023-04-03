package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Set("data", "test")
		c.Next()

		latency := time.Since(t)
		log.Print(latency)

		status := c.Writer.Status()
		log.Print(status)
	}
}
