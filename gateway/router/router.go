package router

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/incwadi-warehouse/monorepo-go/gateway/proxy"
)

func Router() {
	r := gin.New()
	r.SetTrustedProxies(nil)

	if os.Getenv("ENV") != "prod" {
		r.Use(gin.Logger())
	}

	r.Use(gin.Recovery())
	r.Use(headers())

	r.Any("/1/:path", func(c *gin.Context) {
		path := c.Param("path")
		if err := proxy.Proxy(c, os.Getenv("API_CORE"), path); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "Internal Error"})
			return
		}
	})

	r.Run(":8080")
}
