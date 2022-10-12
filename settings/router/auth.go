package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	log.SetPrefix("auth: ")
}

func checkAuth(c *gin.Context) {
	if !isAuthenticated(c.GetHeader("Authorization")) {
	    c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func isAuthenticated(auth string) bool {
	return true
}
