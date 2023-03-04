package web

import (
	"embed"
	"errors"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/incwadi-warehouse/monorepo-go/conf-api/user"
	"github.com/incwadi-warehouse/monorepo-go/conf-api/validation"
)

type Config struct {
	Value string `json:"value"`
}

var schemaName string

var databaseId string

//go:embed data/*
var fs embed.FS

func setSchemaName(d string) error {
	if _, err := fs.ReadFile("data/" + d + ".schema.json"); err != nil {
		return err
	}

	schemaName = d

	return nil
}

func setDatabaseId(d string) error {
	if _, err := fs.ReadFile("data/" + schemaName + ".schema.json"); err != nil {
		return err
	}

	if err := validation.Validate(d, "required,settingsDatabaseId"); err != nil {
		return errors.New("VALIDATION FAILED")
	}

	databaseId = d

	return nil
}

func getSchemaUrl() string {
	return "data/" + schemaName + ".schema.json"
}

func getDefaultsUrl() string {
	return "data/" + schemaName + ".defaults.json"
}

func getDatabaseUrl() string {
	return os.Getenv("DATA_DIR") + schemaName + "-" + databaseId + ".json"
}

func validateDatabaseId(c *gin.Context, schema, id string) bool {
	s := strings.Split(c.GetHeader("Authorization"), " ")
	token := s[1]

	return user.IsTokenValid(token)
}
