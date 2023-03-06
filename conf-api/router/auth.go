package router

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/incwadi-warehouse/monorepo-go/conf-api/user"
)

func checkAuth(c *gin.Context) {
	s := strings.Split(c.GetHeader("Authorization"), " ")

	if hasAuthHeader(s) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if !isAuthenticated(s[1]) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

func hasAuthHeader(s []string) bool {
	return len(s) != 2
}

func isAuthenticated(token string) bool {
    return user.IsTokenValid(token)
}
