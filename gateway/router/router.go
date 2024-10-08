package router

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/incwadi-warehouse/monorepo-go/framework/cors"
	"github.com/incwadi-warehouse/monorepo-go/gateway/proxy"
)

func Router() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.SetTrustedProxies(nil)

    corsConfig := cors.NewCors()
    corsConfig.AllowOrigins = []string{os.Getenv("CORS_ALLOW_ORIGIN")}
    corsConfig.SetCorsHeaders()
	r.Use(corsConfig.SetCorsHeaders())

	r.Any(`/apis/core/1/*path`, func(c *gin.Context) {
		path := c.Param("path")

		safePath := filepath.Join("/", path)

		if err := proxy.Proxy(c, os.Getenv("API_CORE"), safePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "Internal Error"})
			return
		}
	})

	r.Run(":8080")
}
