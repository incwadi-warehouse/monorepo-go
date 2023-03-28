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
	Branch   Branch
}

type Branch struct {
	Id int `json:"id"`
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
	client := &http.Client{}

	req, err := http.NewRequest("GET", os.Getenv("AUTH_API_ME"), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
