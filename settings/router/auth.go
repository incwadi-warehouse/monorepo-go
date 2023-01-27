package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func checkAuth(c *gin.Context) {
	if !isAuthenticated(c.GetHeader("Authorization")) {
	    c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func isAuthenticated(auth string) bool {
	return true
}
