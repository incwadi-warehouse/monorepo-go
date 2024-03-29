package validation

import (
	"os"
	"slices"
	"strings"

	"github.com/go-playground/validator/v10"
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
	var allowedNames []string

    allowedBranches := strings.Split(os.Getenv("BRANCHES"), ",")
	allowedIndexes :=  strings.Split(os.Getenv("INDEXES"), ",")

    for _, name := range allowedIndexes {
		for _, branchId := range allowedBranches {
			allowedNames = append(allowedNames, name+"_"+branchId)
		}
	}

	return slices.Contains(allowedNames, str)
}
