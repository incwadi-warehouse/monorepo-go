package security

import (
	"net/http"
	"testing"
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

var GetDoFunc func(req *http.Request) (*http.Response, error)

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return GetDoFunc(req)
}

func init() {
	Client = &MockClient{}
}

func TestIsTokenValid(t *testing.T) {
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 200}, nil
	}

	if v := IsTokenValid("token"); v == false {
		t.Fatal("Token not valid")
	}
}

func TestUser_HasRole(t *testing.T) {
	type fields struct {
		Id       int
		Username string
		Branch   Branch
		Roles    []string
	}
	type args struct {
		role string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
            "Test valid token",
            fields{
                1,"admin", Branch{1},[]string{"ROLE_ADMIN", "ROLE_USER"},
            },
            args{"ROLE_ADMIN"},
            true,
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := User{
				Id:       tt.fields.Id,
				Username: tt.fields.Username,
				Branch:   tt.fields.Branch,
				Roles:    tt.fields.Roles,
			}
			if got := user.HasRole(tt.args.role); got != tt.want {
				t.Errorf("User.HasRole() = %v, want %v", got, tt.want)
			}
		})
	}
}
