package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/incwadi-warehouse/monorepo-go/gateway/router"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env.local")
	godotenv.Load() // .env

	gin.SetMode(getGinMode())

	router.Router()
}

func getGinMode() string {
	mode := os.Getenv("ENV")

	switch mode {
	case "prod":
		return gin.ReleaseMode
	case "dev":
		return gin.DebugMode
	case "test":
		return gin.TestMode
	default:
		return gin.ReleaseMode
	}
}
