package web

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/incwadi-warehouse/monorepo-go/conf-api/user"
	"github.com/incwadi-warehouse/monorepo-go/conf-api/validation"
)

func Show(c *gin.Context) {
	setSchemaName(c.Param("schemaName"))
	setDatabaseId(c.Param("databaseId"))

	if err := validateParams(c.GetHeader("Authorization")); err != nil {
		c.AbortWithStatus(400)
		return
	}

	s, err := loadDataAndMerge()
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

	if err := validateParams(c.GetHeader("Authorization")); err != nil {
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

	if err := validation.Var(c.Param("key"), "required,confKey"); err != nil {
		c.AbortWithStatus(400)
		return
	}

	s.Add(c.Param("key"), config.Value)

	if err := s.ValidateSchema(); err != nil {
		c.AbortWithStatus(400)
		return
	}

	if err := writeData(s.Data); err != nil {
		c.AbortWithStatus(500)
		return
	}

	d := Response{200, "UPDATED"}

	c.JSON(200, d)
}

func Delete(c *gin.Context) {
	setSchemaName(c.Param("schemaName"))
	setDatabaseId(c.Param("databaseId"))

	if err := validateParams(c.GetHeader("Authorization")); err != nil {
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
		c.AbortWithStatus(500)
		return
	}

	d := Response{200, "SUCCESS"}

	c.JSON(200, d)
}

func validateParams(auth string) error {
    if err := validation.Var(schemaName, "required,confSchemaName"); err != nil {
        return errors.New("INVALID SCHEMA NAME")
    }

	if _, err := fs.ReadFile("data/" + schemaName + ".schema.json"); err != nil {
		return err
	}

    s := strings.Split(auth, " ")
	token := s[1]
	if valid := user.IsTokenValid(token); !valid {
		return errors.New("INVALID DATABASE ID")
	}

	return nil
}
