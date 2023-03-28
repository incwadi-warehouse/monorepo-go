package validation

import (
	"github.com/incwadi-warehouse/monorepo-go/conf-api/user"
	"testing"
)

func TestValidateConfKey(t *testing.T) {
	if err := Var("test", "confKey"); err != nil {
		t.Fatal(err)
	}

	if err := Var("test%", "confKey"); err == nil {
		t.Fatal("Validation must fail")
	}
}

func TestValidateSchemaName(t *testing.T) {
	if err := Var("test", "confSchemaName"); err != nil {
		t.Fatal(err)
	}

	if err := Var("test%", "confSchemaName"); err == nil {
		t.Fatal("Validation must fail")
	}
}

func TestValidateSchemaExists(t *testing.T) {
	if err := Var("user", "confSchemaExists"); err != nil {
		t.Fatal(err)
	}

	if err := Var("test", "confSchemaExists"); err == nil {
		t.Fatal("Validation must fail")
	}
}

func TestValidateDatabaseId(t *testing.T) {
	isTokenValid = func(token string) bool {
		return true
	}

	getUser = func(token string) (user.User, error) {
		return user.User{Id: 1, Username: "admin", Branch: user.Branch{Id: 1}}, nil
	}

	params := Params{
		Auth:       "Bearer token",
		SchemaName: "user",
		DatabaseId: "1",
	}

	if err := Struct(params); err != nil {
		t.Fatal(err)
	}

	params.DatabaseId = "2"

	if err := Struct(params); err == nil {
		t.Fatal("Validation must fail")
	}
}
