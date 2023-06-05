package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/incwadi-warehouse/monorepo-go/search-api/util"
)

var validate *validator.Validate

func init() {
	validate = validator.New()

	validate.RegisterValidation("indexName", validateIndexName)
}

func Var(name interface{}, constraints string) error {
	return validate.Var(name, constraints)
}

func Struct(s interface{}) error {
	return validate.Struct(s)
}

func validateIndexName(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	allowedNames := []string{"products"}

	return util.Contains(str, allowedNames)
}
