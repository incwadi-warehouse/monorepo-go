package web

import (
	"embed"
	"errors"
	"os"
	"strings"

	"github.com/incwadi-warehouse/monorepo-go/conf-api/user"
	"github.com/incwadi-warehouse/monorepo-go/conf-api/validation"
)

type Config struct {
	Value interface{} `json:"value"`
}

type Response struct {
    Status int `json:"status"`
    Message string `json:"message"`
}

var schemaName string

var databaseId string

//go:embed data/*
var fs embed.FS

func setSchemaName(d string) {
	schemaName = d
}

func setDatabaseId(d string) {
	databaseId = d
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

func validateDatabaseId(auth, schema, id string) bool {
	s := strings.Split(auth, " ")
	token := s[1]

	return user.IsTokenValid(token)
}

func validateParams(auth string) error {
	// schema name
	if _, err := fs.ReadFile("data/" + schemaName + ".schema.json"); err != nil {
		return err
	}

	// database id
	if _, err := fs.ReadFile("data/" + schemaName + ".schema.json"); err != nil {
		return err
	}

	if err := validation.Validate(schemaName, "required,settingsDatabaseId"); err != nil {
		return errors.New("VALIDATION FAILED")
	}

    // validate database id
    if valid := validateDatabaseId(auth, schemaName, databaseId); !valid {
        return errors.New("VALIDATION FAILED")
    }

	return nil
}
