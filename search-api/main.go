package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/incwadi-warehouse/monorepo-go/search-api/router"
	"github.com/incwadi-warehouse/monorepo-go/search-api/update"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env.local")
	godotenv.Load() // .env

	update.Run()

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
