package web

import (
	"github.com/gin-gonic/gin"
	"github.com/incwadi-warehouse/monorepo-go/conf-api/validation"
)

func Show(c *gin.Context) {
	setSchemaName(c.Param("schemaName"))
	setDatabaseId(c.Param("databaseId"))

	if err := validate(c.GetHeader("Authorization")); err != nil {
		c.AbortWithStatus(400)
		return
	}

	s, err := loadAndMerge()
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	d := Config{s.Get(c.Param("key"))}

	c.JSON(200, d)
}

func Update(c *gin.Context) {
	setSchemaName(c.Param("schemaName"))
	setDatabaseId(c.Param("databaseId"))

	if err := validate(c.GetHeader("Authorization")); err != nil {
		c.AbortWithStatus(400)
		return
	}

	s, err := loadData()
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	var config Config
	if err := c.ShouldBind(&config); err != nil {
		c.AbortWithStatus(400)
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

	d := Response{200, "UPDATED"}

	c.JSON(200, d)
}

func Delete(c *gin.Context) {
	setSchemaName(c.Param("schemaName"))
	setDatabaseId(c.Param("databaseId"))

	if err := validate(c.GetHeader("Authorization")); err != nil {
		c.AbortWithStatus(400)
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

	d := Response{200,"SUCCESS"}

	c.JSON(200, d)
}
