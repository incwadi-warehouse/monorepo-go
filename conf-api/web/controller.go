package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/incwadi-warehouse/monorepo-go/conf-api/validation"
)

func Show(c *gin.Context) {
	if err := setSchemaName(c.Param("schemaName")); err != nil {
		c.AbortWithStatus(404)
		return
	}

	if err := setDatabaseId(c.Param("databaseId")); err != nil {
		c.AbortWithStatus(404)
		return
	}

	s, err := loadData()
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	data, err := s.Merge()
	if err != nil {
		c.AbortWithStatus(404)
		return
	}
	s.Data = data

	d := Config{fmt.Sprintf("%v", s.Get(c.Param("key")))}

	c.JSON(200, d)
}

func Update(c *gin.Context) {
	if err := setSchemaName(c.Param("schemaName")); err != nil {
		c.AbortWithStatus(404)
		return
	}

	if err := setDatabaseId(c.Param("databaseId")); err != nil {
		c.AbortWithStatus(404)
		return
	}

	s, err := loadData()
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	var config Config
	if err := c.ShouldBind(&config); err != nil {
		c.AbortWithStatus(404)
		return
	}

	if err := validation.Validate(c.Param("key"), "required,settingsKey"); err != nil {
		c.AbortWithStatus(400)
		return
	}

	s.Add(c.Param("key"), config.Value)

	if err := s.ValidateSchema(); err != nil {
		c.AbortWithStatus(400)
		return
	}

	if err := writeData(s.Data); err != nil {
		c.AbortWithStatus(404)
		return
	}

	d := Config{c.Param("key")}

	c.JSON(200, d)
}

func Delete(c *gin.Context) {
	if err := setSchemaName(c.Param("schemaName")); err != nil {
		c.AbortWithStatus(404)
		return
	}

	if err := setDatabaseId(c.Param("databaseId")); err != nil {
		c.AbortWithStatus(404)
		return
	}

	s, err := loadData()
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	s.Rm(c.Param("key"))

	if err := s.ValidateSchema(); err != nil {
		c.AbortWithStatus(400)
		return
	}

	if err := writeData(s.Data); err != nil {
		c.AbortWithStatus(404)
		return
	}

	d := Config{"SUCCESS"}

	c.JSON(200, d)
}
