package router

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/incwadi-warehouse/monorepo-go/blog/content/article"
	"github.com/incwadi-warehouse/monorepo-go/blog/content/home"
	"github.com/incwadi-warehouse/monorepo-go/framework/apikey"
	"github.com/incwadi-warehouse/monorepo-go/framework/cors"
	"github.com/spf13/viper"
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

// permissionsMiddleware is a middleware to check for API key permissions.
func permissionsMiddleware(permissions ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("X-API-Key")

		for _, permission := range permissions {
			if !apikey.HasPermission(key, permission) {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
				return
			}
		}

		c.Next()
	}
}

// setupRouter sets up the Gin router.
func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.Use(cors.SetCorsHeaders(viper.GetString("CORS_ALLOW_ORIGIN")))

	api := r.Group("/", authMiddleware)
	{
		api.GET("/home", permissionsMiddleware("articles"), func(c *gin.Context) {
			index, err := home.GetHome()
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.String(http.StatusOK, index)
		})
		api.GET("/home/new/:days", permissionsMiddleware("articles"), func(c *gin.Context) {
			daysStr := c.Param("days")

			days, err := strconv.Atoi(daysStr)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid 'days' parameter"})
				return
			}

			newCount, err := home.GetNewArticles(days)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"new_articles": newCount})
		})
		api.GET("/article/*path", permissionsMiddleware("articles"), func(c *gin.Context) {
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
