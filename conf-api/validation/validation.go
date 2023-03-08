package validation

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/incwadi-warehouse/monorepo-go/conf-api/schema"
	"github.com/incwadi-warehouse/monorepo-go/conf-api/user"
)

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

func validateDatabaseId(fl validator.FieldLevel) bool {
	f := fl.Field().Interface().(map[string]interface{})

	s := strings.Split(f["auth"].(string), " ")
	token := s[1]

	if valid := user.IsTokenValid(token); !valid {
		return false
	}

	u, _ := user.GetUser(token)

	if f["schemaName"] == "user" {
		return strconv.Itoa(u.Id) == f["databaseId"].(string)
	}

	if f["schemaName"] == "branch" {
		return strconv.Itoa(u.Branch.Id) == f["databaseId"].(string)
	}

	return false
}
