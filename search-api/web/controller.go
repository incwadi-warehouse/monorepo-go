package web

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/incwadi-warehouse/monorepo-go/search-api/api"
	"github.com/incwadi-warehouse/monorepo-go/search-api/util"
	"github.com/incwadi-warehouse/monorepo-go/search-api/validation"
	"github.com/incwadi-warehouse/monorepo-go/security/authentication"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func Search(c *gin.Context) {
	if validation.Var(c.Param("index"), "indexName") != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, Response{http.StatusNotFound, "Not Found"})
		return
	}

	status, data := api.NewRequest("POST", "/indexes/"+c.Param("index")+"/search", c.Request.Body)

	c.JSON(status, data)
}

func Rebuild(c *gin.Context) {
	auth, exists := c.Get("auth")
	if !exists {
		c.AbortWithStatusJSON(http.StatusForbidden, Response{http.StatusForbidden, "Forbidden"})
		return
	}

	if !util.Contains("ROLE_ADMIN", auth.(authentication.Auth).User.Roles) {
		c.AbortWithStatusJSON(http.StatusForbidden, Response{http.StatusForbidden, "Forbidden"})
		return
	}

	if validation.Var(c.Param("index"), "indexName") != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, Response{http.StatusNotFound, "Not Found"})
		return
	}

	api.NewRequest("DELETE", "/indexes/"+c.Param("index")+"/documents", strings.NewReader(""))
	status, data := api.NewRequest("POST", "/indexes/"+c.Param("index")+"/documents", c.Request.Body)

	c.JSON(status, data)
}
