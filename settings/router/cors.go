package router

import (
	"log"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	log.SetPrefix("cors: ")
}

func headers() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:  strings.Split(os.Getenv("CORS_ALLOW_ORIGIN"), ","),
		AllowMethods:  []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:  []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
	})
}
