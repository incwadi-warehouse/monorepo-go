package authentication

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
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

func TestGetUser(t *testing.T) {
	user, err := json.Marshal(
		User{1, "admin", Branch{1}, []string{"ROLE_USER"}},
	)
	if err != nil {
		t.Fatal(err)
	}

	GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(string(user))),
		}, nil
	}

	auth, err := GetUser("token")
	if err != nil {
		t.Fatal("No auth received", auth)
	}
}
