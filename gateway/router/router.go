package router

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/incwadi-warehouse/monorepo-go/cors/cors"
	"github.com/incwadi-warehouse/monorepo-go/gateway/proxy"
)

func Router() {
    gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.Use(cors.Headers())

	r.Any(`/apis/core/1/*path`, func(c *gin.Context) {
		path := c.Param("path")
		if err := proxy.Proxy(c, os.Getenv("API_CORE"), path); err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"msg": "Internal Error"})
			return
		}
	})

	r.Run(":8080")
}
