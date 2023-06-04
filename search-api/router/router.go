package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/incwadi-warehouse/monorepo-go/search-api/mock"
	"github.com/incwadi-warehouse/monorepo-go/search-api/web"
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

	auth.POST("/indexes/:index/search", web.Search)
	auth.GET("/indexes", web.List)
	auth.POST("/indexes", web.Create)
	auth.DELETE("/indexes/:index", web.Remove)

	auth.GET("/indexes/:index/settings", web.GetSettings)
	auth.PATCH("/indexes/:index/settings", web.UpdateSettings)

	auth.DELETE("/indexes/:index/documents", web.RemoveDocuments)
	auth.POST("/indexes/:index/documents", web.CreateDocument)

    if os.Getenv("ENV") != "prod" {
        r.GET("/api/me", mock.Me)
    }

	r.Run(":8080")
}
