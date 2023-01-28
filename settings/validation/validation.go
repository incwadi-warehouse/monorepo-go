package validation

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()

	validate.RegisterValidation("settingsKey", validateSettingsKey)
}

func Validate(name string) error {
	return validate.Var(name, "required,settingsKey")
}

func validateSettingsKey(fl validator.FieldLevel) bool {
	s := strings.Split(fl.Field().String(), ".")
	re := regexp.MustCompile(`^[a-zA-Z0-9-_]+$`)

	for _, v := range s {
		if matched := re.MatchString(v); !matched {
			return false
		}
	}

	return true
}
