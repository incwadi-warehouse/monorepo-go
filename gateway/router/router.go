package router

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/incwadi-warehouse/monorepo-go/gateway/proxy"
)

func Routes() *gin.Engine {
	r := gin.Default()

	r.Any(`/apis/core/1/*path`, func(c *gin.Context) {
		path := c.Param("path")

		safePath := filepath.Join("/", path)

		if err := proxy.Proxy(c, os.Getenv("API_CORE"), safePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "Internal Error"})
			return
		}
	})

	return r
}
