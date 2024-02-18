package middleware

import (
	"bytes"
	"io"
	"log"
	"time"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func HttpRequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Println("Error reading request body:", err)
			c.Next()
			return
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		start := time.Now()

		c.Next()

		end := time.Now()

		color.Cyan("Request Body: %s\n", body)
		color.Cyan("Request processed in %v\n", end.Sub(start))
	}
}
