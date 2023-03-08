package validation

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()

	validate.RegisterValidation("confKey", validateConfKey)
	validate.RegisterValidation("confSchemaName", validateSchemaName)
}

func Var(name, constraints string) error {
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
