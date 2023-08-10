package web

import (
	"encoding/json"
	"io"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/incwadi-warehouse/monorepo-go/search-api/meili"
	"github.com/incwadi-warehouse/monorepo-go/search-api/validation"
	"github.com/incwadi-warehouse/monorepo-go/security/authentication"
	"github.com/meilisearch/meilisearch-go"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SearchQuery struct {
	Q      string      `json:"q"`
	Limit  int64       `json:"limit"`
	Filter interface{} `json:"filter"`
	Facets []string    `json:"facets"`
}

func Search(c *gin.Context) {
	if validation.Var(c.Param("index"), "indexName") != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, Response{http.StatusNotFound, "Not Found"})
		return
	}

	io, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, "Server Error"})
		return
	}

	var s SearchQuery
	if err := json.Unmarshal(io, &s); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, "Server Error"})
		return
	}

	client := meili.NewClient()
	data, err := client.Index(c.Param("index")).Search(s.Q, &meilisearch.SearchRequest{
		Limit:  s.Limit,
		Filter: s.Filter,
		Facets: s.Facets,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, "Server Error"})
		return
	}

	c.JSON(http.StatusOK, data)
}

func Rebuild(c *gin.Context) {
	auth, exists := c.Get("auth")
	if !exists {
		c.AbortWithStatusJSON(http.StatusForbidden, Response{http.StatusForbidden, "Forbidden"})
		return
	}

	if !slices.Contains(auth.(authentication.Auth).User.Roles, "ROLE_ADMIN") {
		c.AbortWithStatusJSON(http.StatusForbidden, Response{http.StatusForbidden, "Forbidden"})
		return
	}

	if validation.Var(c.Param("index"), "indexName") != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, Response{http.StatusNotFound, "Not Found"})
		return
	}

	client := meili.NewClient()
	client.Index(c.Param("index")).DeleteAllDocuments()
	client.Index(c.Param("index")).UpdateDocuments(c.Request.Body)

	c.JSON(http.StatusOK, Response{http.StatusOK, "Documents added to queue."})
}
