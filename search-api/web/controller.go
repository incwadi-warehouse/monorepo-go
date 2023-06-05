package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/incwadi-warehouse/monorepo-go/search-api/api"
	"github.com/incwadi-warehouse/monorepo-go/search-api/validation"
)

func Search(c *gin.Context) {
	if validation.Var(c.Param("index"), "indexName") != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, Response{http.StatusNotFound, "Not Found"})
		return
	}

	var body SearchBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{http.StatusBadRequest, "No Valid Data"})
		return
	}
	if validation.Struct(body) != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{http.StatusBadRequest, "Not Valid"})
		return
	}

	status, data := api.NewRequest("POST", "/indexes/"+c.Param("index")+"/search", body)

	c.JSON(status, data)
}

func List(c *gin.Context) {
	status, data := api.NewRequest("GET", "/indexes", nil)

	c.JSON(status, data)
}

func Create(c *gin.Context) {
	var body IndexBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{http.StatusBadRequest, "No Valid Data"})
		return
	}
	if validation.Struct(body) != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{http.StatusBadRequest, "Not Valid"})
		return
	}

	status, data := api.NewRequest("POST", "/indexes", body)

	c.JSON(status, data)
}

func Remove(c *gin.Context) {
	// if validation.Var(c.Param("index"), "indexName") != nil {
	// 	c.AbortWithStatusJSON(http.StatusNotFound, Response{http.StatusNotFound, "Not Found"})
	// 	return
	// }

	status, data := api.NewRequest("DELETE", "/indexes/"+c.Param("index"), nil)

	c.JSON(status, data)
}

func RemoveDocuments(c *gin.Context) {
	if validation.Var(c.Param("index"), "indexName") != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, Response{http.StatusNotFound, "Not Found"})
		return
	}

	status, data := api.NewRequest("DELETE", "/indexes/"+c.Param("index")+"/documents", nil)

	c.JSON(status, data)
}

func CreateDocument(c *gin.Context) {
	if validation.Var(c.Param("index"), "indexName") != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, Response{http.StatusNotFound, "Not Found"})
		return
	}

	var body []DocumentsBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{http.StatusBadRequest, "No Valid Data"})
		return
	}

    // fmt.Print(validation.Struct(body))
	// if validation.Struct(body) != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, Response{http.StatusBadRequest, "Not Valid"})
	// 	return
	// }

	status, data := api.NewRequest("POST", "/indexes/"+c.Param("index")+"/documents", body)

	c.JSON(status, data)
}

func GetSettings(c *gin.Context) {
	if validation.Var(c.Param("index"), "indexName") != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, Response{http.StatusNotFound, "Not Found"})
		return
	}

	status, data := api.NewRequest("GET", "/indexes/"+c.Param("index")+"/settings", nil)

	c.JSON(status, data)
}

func UpdateSettings(c *gin.Context) {
	if validation.Var(c.Param("index"), "indexName") != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, Response{http.StatusNotFound, "Not Found"})
		return
	}

	var body SettingsBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{http.StatusBadRequest, "No Valid Data"})
		return
	}
	if validation.Struct(body) != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{http.StatusBadRequest, "Not Valid"})
		return
	}

	status, data := api.NewRequest("PATCH", "/indexes/"+c.Param("index")+"/settings", body)

	c.JSON(status, data)
}
