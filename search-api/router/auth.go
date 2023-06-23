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

	if len(s) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, Response{401, "Token missing"})
		return
	}

	auth,_ := security.GetUser(s[1])

	if !auth.IsAuthenticated {
		c.AbortWithStatusJSON(http.StatusUnauthorized, Response{401, "Token missing"})
		return
	}

	c.Set("auth", auth)
}
