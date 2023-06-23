package router

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/incwadi-warehouse/monorepo-go/security/auth"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func checkAuth(c *gin.Context) {
	s := strings.Split(c.GetHeader("Authorization"), " ")

	if len(s) != 2 {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			Response{http.StatusUnauthorized, "Token missing"},
		)
		return
	}

	auth, err := auth.GetUser(s[1])
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			Response{http.StatusUnauthorized, "Unauthorized"},
		)
		return
	}

	if !auth.IsAuthenticated {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			Response{http.StatusUnauthorized, "Unauthorized"},
		)
		return
	}

	c.Set("auth", auth)
}
