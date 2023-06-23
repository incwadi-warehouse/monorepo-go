package security

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
)

type Auth struct {
	User            User
	IsAuthenticated bool
}

type User struct {
	Id       int      `json:"id"`
	Username string   `json:"username"`
	Branch   Branch   `json:"branch"`
	Roles    []string `json:"roles"`
}

type Branch struct {
	Id int `json:"id"`
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var Client HTTPClient

func init() {
	Client = &http.Client{}
}

func GetUser(token string) (Auth, error) {
	var auth Auth = Auth{
		IsAuthenticated: false,
	}

	res, err := request(token)
	if err != nil {
		return auth, err
	}

	io, err := io.ReadAll(res.Body)
	if err != nil {
		return auth, err
	}
	if err := json.Unmarshal(io, &auth.User); err != nil {
		return auth, err
	}

	auth.IsAuthenticated = true

	return auth, nil
}

func request(token string) (*http.Response, error) {
	req, err := http.NewRequest("GET", os.Getenv("AUTH_API_ME"), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)

	res, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New("TOKEN NOT VALID")
	}

	return res, nil
}
