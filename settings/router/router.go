package router

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/incwadi-warehouse/monorepo-go/settings/branch"
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

	// branch
	auth.GET("/:databaseName/:key", branch.Show)
	auth.POST("/:databaseName/:key", branch.Update)
	auth.PUT("/:databaseName/:key", branch.Update)
	auth.DELETE("/:databaseName/:key", branch.Delete)

	r.Run(":8080")
}
