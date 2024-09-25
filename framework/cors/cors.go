package cors

import (
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Headers(allowedOrigins string) gin.HandlerFunc {
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
