package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/incwadi-warehouse/monorepo-go/search-api/api"
	"github.com/incwadi-warehouse/monorepo-go/search-api/validation"
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

func List(c *gin.Context) {
	status, data := api.NewRequest("GET", "/indexes", c.Request.Body)

	c.JSON(status, data)
}

func Create(c *gin.Context) {

	status, data := api.NewRequest("POST", "/indexes", c.Request.Body)

	c.JSON(status, data)
}

func Remove(c *gin.Context) {
	// if validation.Var(c.Param("index"), "indexName") != nil {
	// 	c.AbortWithStatusJSON(http.StatusNotFound, Response{http.StatusNotFound, "Not Found"})
	// 	return
	// }

	status, data := api.NewRequest("DELETE", "/indexes/"+c.Param("index"), c.Request.Body)

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
	if validation.Var(c.Param("index"), "indexName") != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, Response{http.StatusNotFound, "Not Found"})
		return
	}

	status, data := api.NewRequest("GET", "/indexes/"+c.Param("index")+"/settings", c.Request.Body)

	c.JSON(status, data)
}

func UpdateSettings(c *gin.Context) {
	if validation.Var(c.Param("index"), "indexName") != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, Response{http.StatusNotFound, "Not Found"})
		return
	}

	status, data := api.NewRequest("PATCH", "/indexes/"+c.Param("index")+"/settings", c.Request.Body)

	c.JSON(status, data)
}
