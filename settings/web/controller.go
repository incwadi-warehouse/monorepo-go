package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Show(c *gin.Context) {
    if err := setDatabaseName(c.Param("databaseName")); err != nil {
        c.AbortWithStatus(404)
        return
    }

	s, err := loadData()
	if err != nil {
		c.AbortWithStatus(404)
        return
	}

	d := Config{fmt.Sprintf("%v", s.Get(c.Param("key")))}

	c.JSON(200, d)
}

func Update(c *gin.Context) {
    if err := setDatabaseName(c.Param("databaseName")); err != nil {
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

	s.Add(c.Param("key"), config.Value)

	if err := writeData(s.Data); err != nil {
        c.AbortWithStatus(404)
        return
    }

	d := Config{c.Param("key")}

	c.JSON(200, d)
}

func Delete(c *gin.Context) {
    if err := setDatabaseName(c.Param("databaseName")); err != nil {
        c.AbortWithStatus(404)
        return
    }

	s, err := loadData()
	if err != nil {
		c.AbortWithStatus(404)
        return
	}

	s.Rm(c.Param("key"))

    if err := writeData(s.Data); err != nil {
        c.AbortWithStatus(404)
        return
    }

	d := Config{"SUCCESS"}

	c.JSON(200, d)
}
