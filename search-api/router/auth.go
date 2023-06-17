package router

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/incwadi-warehouse/monorepo-go/security/security"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func checkAuth(c *gin.Context) {
	s := strings.Split(c.GetHeader("Authorization"), " ")

	if !hasAuthHeader(s) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, Response{401, "Token missing"})
		return
	}

	if !isAuthenticated(s[1]) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, Response{401, "Token missing"})
		return
	}
}

func hasAuthHeader(s []string) bool {
	return len(s) == 2
}

func isAuthenticated(token string) bool {
	return security.IsTokenValid(token)
}
