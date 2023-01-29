package router

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func checkAuth(c *gin.Context) {
	s := strings.Split(c.GetHeader("Authorization"), " ")

	if hasAuthHeader(s) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	token := s[1]

	if !isAuthenticated(token) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

func hasAuthHeader(s []string) bool {
	return len(s) != 2
}

func isAuthenticated(auth string) bool {
	return auth == os.Getenv("API_KEY")
}
