package router

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/incwadi-warehouse/monorepo-go/settings/web"
)

func init() {
	log.SetPrefix("router: ")
}

func Router() {
	r := gin.New()
	r.SetTrustedProxies(nil)

	if os.Getenv("ENV") != "prod" {
		r.Use(gin.Logger())
	}

	r.Use(gin.Recovery())

	r.Use(headers())

	auth := r.Group("/api", checkAuth)

	// web
	auth.GET("/:databaseName/:key", web.Show)
	auth.POST("/:databaseName/:key", web.Update)
	auth.PUT("/:databaseName/:key", web.Update)
	auth.DELETE("/:databaseName/:key", web.Delete)

	r.Run(":8080")
}
