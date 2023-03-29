package user

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
