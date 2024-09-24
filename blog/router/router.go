package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/incwadi-warehouse/monorepo-go/blog/apikey"
	"github.com/incwadi-warehouse/monorepo-go/blog/content/article"
	"github.com/incwadi-warehouse/monorepo-go/blog/content/home"
	"github.com/incwadi-warehouse/monorepo-go/cors/cors"
)

// authMiddleware is a middleware to check for API key authentication.
func authMiddleware(c *gin.Context) {
	key := c.GetHeader("X-API-Key")

	if !apikey.IsValidAPIKey(key) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
		return
	}

	c.Next()
}

// setupRouter sets up the Gin router.
func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.Use(cors.Headers())

	api := r.Group("/", authMiddleware)
	{
		api.GET("/home", func(c *gin.Context) {
			index, err := home.GetHome()
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.String(http.StatusOK, index)
		})
		api.GET("/article/*path", func(c *gin.Context) {
			path := c.Param("path")

			cnt, err := article.GetArticle(path)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.String(http.StatusOK, cnt)
		})
	}

	return r
}
