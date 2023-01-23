package branch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/incwadi-warehouse/monorepo-go/conf/settings"
	"github.com/incwadi-warehouse/monorepo-go/settings/storage"
)

func init() {
	log.SetPrefix("branch: ")

	if err := storage.Exists(getDatabaseUrl()); err != nil {
		writeBaseConfig()
	}
}

func Show(c *gin.Context) {
    schema, err := os.ReadFile(getSchemaUrl())
    if err != nil {
        c.AbortWithStatus(404)
    }

    defaults, err := os.ReadFile(getDefaultsUrl())
    if err != nil {
        c.AbortWithStatus(404)
    }

    file, err1 := os.ReadFile(getDatabaseUrl())
    if err1 != nil {
        c.AbortWithStatus(404)
    }

	s, err := settings.LoadFromString(schema, defaults, file)
	if err != nil {
		c.AbortWithStatus(404)
	}

	d := Config{fmt.Sprintf("%v", s.Get(c.Param("key")))}

	c.JSON(200, d)
}

func Update(c *gin.Context) {
    schema, err := os.ReadFile(getSchemaUrl())
    if err != nil {
        c.AbortWithStatus(404)
    }

    defaults, err := os.ReadFile(getDefaultsUrl())
    if err != nil {
        c.AbortWithStatus(404)
    }

    file, err1 := os.ReadFile(getDatabaseUrl())
    if err1 != nil {
        c.AbortWithStatus(404)
    }

	s, err := settings.LoadFromString(schema, defaults, file)
	if err != nil {
		c.AbortWithStatus(404)
	}

	var config Config
	if err := c.ShouldBind(&config); err != nil {
		c.AbortWithStatus(404)
	}

	s.Add(c.Param("key"), config.Value)
	v, err := json.Marshal(s.Value)
	if err != nil {
		c.AbortWithStatus(404)
	}

	var out bytes.Buffer
	if err := json.Indent(&out, v, "", "\t"); err != nil {
		c.AbortWithStatus(404)
	}

    if err := storage.Write(getDatabaseUrl(), out.Bytes()); err != nil{
        c.AbortWithStatus(404)
    }

	d := Config{c.Param("key")}

	c.JSON(200, d)
}

func Delete(c *gin.Context) {
    schema, err := os.ReadFile(getSchemaUrl())
    if err != nil {
        c.AbortWithStatus(404)
    }

    defaults, err := os.ReadFile(getDefaultsUrl())
    if err != nil {
        c.AbortWithStatus(404)
    }

    file, err1 := os.ReadFile(getDatabaseUrl())
    if err1 != nil {
        c.AbortWithStatus(404)
    }

	s, err := settings.LoadFromString(schema, defaults, file)
	if err != nil {
		c.AbortWithStatus(404)
	}

	s.Rm(c.Param("key"))
	v, err := json.Marshal(s.Value)
	if err != nil {
		c.AbortWithStatus(404)
	}

	var out bytes.Buffer
	if err := json.Indent(&out, v, "", "\t"); err != nil {
		c.AbortWithStatus(404)
	}

    if err := storage.Write(getDatabaseUrl(), out.Bytes()); err != nil{
        c.AbortWithStatus(404)
    }

	d := Config{"SUCCESS"}

	c.JSON(200, d)
}
