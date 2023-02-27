package middlewares

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")

		if token != requiredToken {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid API token"})
			return
		}

		c.Next()
	}
}

func ResponseMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		url := c.Request.URL
		host := c.Request.Host
		meth := c.Request.Method
		t := time.Now().Format("2006-01-02 15:04:05")
		size := c.Request.ContentLength
		fmt.Println("\nMetodo: ", meth, "\nURL: ", host, url, "\nFecha: ", t, "\nTamaño: ", size, "\n\n ")
		c.Next()
	}
}
