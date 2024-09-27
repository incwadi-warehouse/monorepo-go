package cors

import (
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetCorsHeaders sets up CORS middleware with the given allowed origins.
// If allowedOrigins is empty, it defaults to allowing all origins ("*").
func SetCorsHeaders(allowedOrigins string) gin.HandlerFunc {
	if allowedOrigins == "" {
		allowedOrigins = "*"
	}

	return cors.New(cors.Config{
		AllowOrigins:  strings.Split(allowedOrigins, ","),
		AllowMethods:  []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:  []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
	})
}
