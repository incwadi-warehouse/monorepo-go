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
	validate.RegisterValidation("settingsDatabaseId", validateDatabaseId)
}

func Validate(name, constraints string) error {
	return validate.Var(name, constraints)
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

func validateDatabaseId(fl validator.FieldLevel) bool {
    re := regexp.MustCompile(`^[a-zA-Z0-9-_]+$`)

    return re.MatchString(fl.Field().String())
}
