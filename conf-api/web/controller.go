package web

import (
	"github.com/gin-gonic/gin"
	"github.com/incwadi-warehouse/monorepo-go/conf-api/storage"
	"github.com/incwadi-warehouse/monorepo-go/conf-api/validation"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Config struct {
	Value interface{} `json:"value"`
}

var (
	schemaName string
	databaseId string
)

func Show(c *gin.Context) {
	params := validation.Params{
		Auth:       c.GetHeader("Authorization"),
		SchemaName: c.Param("schemaName"),
		DatabaseId: c.Param("databaseId"),
	}

	if err := validation.Struct(params); err != nil {
		c.AbortWithStatus(400)
		return
	}

	schemaName = c.Param("schemaName")
	databaseId = c.Param("databaseId")

	s, err := storage.LoadDataAndMerge(schemaName, databaseId)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	d := Config{s.Get(c.Param("key"))}

	c.JSON(200, d)
}

func Update(c *gin.Context) {
	params := validation.Params{
		Auth:       c.GetHeader("Authorization"),
		SchemaName: c.Param("schemaName"),
		DatabaseId: c.Param("databaseId"),
	}

	if err := validation.Struct(params); err != nil {
		c.AbortWithStatus(400)
		return
	}

	schemaName = c.Param("schemaName")
	databaseId = c.Param("databaseId")

	s, err := storage.LoadData(schemaName, databaseId)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	var config Config
	if err := c.ShouldBind(&config); err != nil {
		c.AbortWithStatus(400)
		return
	}

	if err := validation.Var(c.Param("key"), "required,confKey"); err != nil {
		c.AbortWithStatus(400)
		return
	}

	s.Add(c.Param("key"), config.Value)

	if err := s.ValidateSchema(); err != nil {
		c.AbortWithStatus(400)
		return
	}

	if err := storage.WriteData(schemaName, databaseId, s.Data); err != nil {
		c.AbortWithStatus(500)
		return
	}

	d := Response{200, "UPDATED"}

	c.JSON(200, d)
}

func Delete(c *gin.Context) {
	params := validation.Params{
		Auth:       c.GetHeader("Authorization"),
		SchemaName: c.Param("schemaName"),
		DatabaseId: c.Param("databaseId"),
	}
	if err := validation.Struct(params); err != nil {
		c.AbortWithStatus(400)
		return
	}

	schemaName = c.Param("schemaName")
	databaseId = c.Param("databaseId")

	s, err := storage.LoadData(schemaName, databaseId)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	s.Rm(c.Param("key"))

	if err := s.ValidateSchema(); err != nil {
		c.AbortWithStatus(400)
		return
	}

	if err := storage.WriteData(schemaName, databaseId, s.Data); err != nil {
		c.AbortWithStatus(500)
		return
	}

	d := Response{200, "SUCCESS"}

	c.JSON(200, d)
}
