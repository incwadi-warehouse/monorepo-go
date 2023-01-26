package branch

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/incwadi-warehouse/monorepo-go/settings/storage"
)

func init() {
	if err := storage.Exists(getDatabaseUrl()); err != nil {
		writeBaseConfig()
	}
}

func Show(c *gin.Context) {
	s, err := loadData()
	if err != nil {
		c.AbortWithStatus(404)
	}

	d := Config{fmt.Sprintf("%v", s.Get(c.Param("key")))}

	c.JSON(200, d)
}

func Update(c *gin.Context) {
	s, err := loadData()
	if err != nil {
		c.AbortWithStatus(404)
	}

	var config Config
	if err := c.ShouldBind(&config); err != nil {
		c.AbortWithStatus(404)
	}

	s.Add(c.Param("key"), config.Value)

	if err := writeData(s.Data); err != nil {
        c.AbortWithStatus(404)
    }

	d := Config{c.Param("key")}

	c.JSON(200, d)
}

func Delete(c *gin.Context) {
	s, err := loadData()
	if err != nil {
		c.AbortWithStatus(404)
	}

	s.Rm(c.Param("key"))

    if err := writeData(s.Data); err != nil {
        c.AbortWithStatus(404)
    }

	d := Config{"SUCCESS"}

	c.JSON(200, d)
}
