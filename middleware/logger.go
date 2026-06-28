package middleware

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {

	file, _ := os.OpenFile(
		"logs/server.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666,
	)

	logger := log.New(file, "", log.LstdFlags)

	return func(c *gin.Context) {

		start := time.Now()

		c.Next()

		logger.Printf(
			"%s | %s | %s | %v",
			c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP(),
			time.Since(start),
		)

	}

}
