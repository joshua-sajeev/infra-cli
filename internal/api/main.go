package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	// defer stop()
	r := gin.Default()

	// Define a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong!",
		})
	})
	r.Run(":8080")

}
