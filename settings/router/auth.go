package router

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func checkAuth(c *gin.Context) {
    s := strings.Split(c.GetHeader("Authorization"), " ")

    if len(s) != 2 {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }
    token := s[1]

	if !isAuthenticated(token) {
	    c.AbortWithStatus(http.StatusUnauthorized)
        return
	}
}

func isAuthenticated(auth string) bool {
	return auth == os.Getenv("API_KEY")
}
