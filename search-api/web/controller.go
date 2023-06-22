package web

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/incwadi-warehouse/monorepo-go/search-api/api"
	"github.com/incwadi-warehouse/monorepo-go/search-api/validation"
	"github.com/incwadi-warehouse/monorepo-go/security/security"
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

func RemoveDocuments(c *gin.Context) {
	if validation.Var(c.Param("index"), "indexName") != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, Response{http.StatusNotFound, "Not Found"})
		return
	}

	status, data := api.NewRequest("DELETE", "/indexes/"+c.Param("index")+"/documents", c.Request.Body)

	c.JSON(status, data)
}

func CreateDocument(c *gin.Context) {
	if validation.Var(c.Param("index"), "indexName") != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, Response{http.StatusNotFound, "Not Found"})
		return
	}

	status, data := api.NewRequest("POST", "/indexes/"+c.Param("index")+"/documents", c.Request.Body)

	c.JSON(status, data)
}

func GetSettings(c *gin.Context) {
	if !hasRole(c.GetHeader("Authorization"), "ROLE_ADMIN") {
		c.AbortWithStatusJSON(http.StatusForbidden, Response{http.StatusForbidden, "Forbidden"})
		return
	}

	if validation.Var(c.Param("index"), "indexName") != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, Response{http.StatusNotFound, "Not Found"})
		return
	}

	status, data := api.NewRequest("GET", "/indexes/"+c.Param("index")+"/settings", c.Request.Body)

	c.JSON(status, data)
}

func UpdateSettings(c *gin.Context) {
	if !hasRole(c.GetHeader("Authorization"), "ROLE_ADMIN") {
		c.AbortWithStatusJSON(http.StatusForbidden, Response{http.StatusForbidden, "Forbidden"})
		return
	}

	if validation.Var(c.Param("index"), "indexName") != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, Response{http.StatusNotFound, "Not Found"})
		return
	}

	status, data := api.NewRequest("PATCH", "/indexes/"+c.Param("index")+"/settings", c.Request.Body)

	c.JSON(status, data)
}

func hasRole (auth string, role string) bool {
    token := strings.Split(auth, " ")
    if len(token) != 2 {
        return false
    }
	user, err := security.GetUser(token[1])
	if err != nil {
		return false
	}

    return user.HasRole(role)
}
