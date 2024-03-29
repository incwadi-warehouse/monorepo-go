package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/incwadi-warehouse/monorepo-go/conf-api/mock"
	"github.com/incwadi-warehouse/monorepo-go/conf-api/web"
)

func Router() {
	r := gin.New()
	r.SetTrustedProxies(nil)

	if os.Getenv("ENV") != "prod" {
		r.Use(gin.Logger())
	}

	r.Use(gin.Recovery())

	r.Use(headers())

	auth := r.Group("/" + os.Getenv("BASE_PATH") + "/api", checkAuth)

	auth.GET("/:schemaName/:databaseId/:key", web.Show)
	auth.POST("/:schemaName/:databaseId/:key", web.Update)
	auth.PUT("/:schemaName/:databaseId/:key", web.Update)
	auth.DELETE("/:schemaName/:databaseId/:key", web.Delete)

    if os.Getenv("ENV") != "prod" {
        r.GET("/api/me", mock.Me)
    }

	r.Run(":8080")
}
