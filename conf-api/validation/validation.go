package validation

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/incwadi-warehouse/monorepo-go/conf-api/schema"
	"github.com/incwadi-warehouse/monorepo-go/conf-api/user"
)

type Params struct {
	Auth       string `validate:"required"`
	SchemaName string `validate:"required,confSchemaName,confSchemaExists"`
	DatabaseId string `validate:"required,confDatabaseId"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()

	validate.RegisterValidation("confKey", validateConfKey)
	validate.RegisterValidation("confSchemaName", validateSchemaName)
	validate.RegisterValidation("confSchemaExists", validateSchemaExists)
	validate.RegisterValidation("confDatabaseId", validateDatabaseId)
}

func Var(name interface{}, constraints string) error {
	return validate.Var(name, constraints)
}

func Struct(s interface{}) error {
	return validate.Struct(s)
}

func validateConfKey(fl validator.FieldLevel) bool {
	s := strings.Split(fl.Field().String(), ".")
	re := regexp.MustCompile(`^[a-zA-Z0-9-_]+$`)

	for _, v := range s {
		if matched := re.MatchString(v); !matched {
			return false
		}
	}

	return true
}

func validateSchemaName(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9-_]+$`)

	return re.MatchString(fl.Field().String())
}

func validateSchemaExists(fl validator.FieldLevel) bool {
	if _, err := schema.Fs.ReadFile("data/" + fl.Field().String() + ".schema.json"); err != nil {
		return false
	}

	return true
}

var isTokenValid = user.IsTokenValid
var getUser = user.GetUser

func validateDatabaseId(fl validator.FieldLevel) bool {
	f := fl.Parent().Interface().(Params)

	s := strings.Split(f.Auth, " ")
	token := s[1]

	if valid := isTokenValid(token); !valid {
		return false
	}

	u, _ := getUser(token)

	if f.SchemaName == "user" {
		return strconv.Itoa(u.Id) == f.DatabaseId
	}

	if f.SchemaName == "branch" {
		return strconv.Itoa(u.Branch.Id) == f.DatabaseId
	}

	return false
}
