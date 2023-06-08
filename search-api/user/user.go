package user

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Branch   Branch `json:"branch"`
	Roles   []string `json:"roles"`
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

func IsTokenValid(token string) bool {
	res, err := request(token)
	if err != nil {
		return false
	}

	return res.StatusCode == 200
}

func GetUser(token string) (User, error) {
	var u User

	res, err := request(token)
	if err != nil {
		return u, err
	}

	io, err := io.ReadAll(res.Body)
	if err != nil {
		return u, err
	}

	if err := json.Unmarshal(io, &u); err != nil {
		return u, err
	}

	return u, nil
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

	return res, nil
}
