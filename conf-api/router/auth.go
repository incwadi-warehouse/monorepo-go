package router

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
    Branch struct {
        Id int `json:"id"`
    }
}

func checkAuth(c *gin.Context) {
	s := strings.Split(c.GetHeader("Authorization"), " ")

	if hasAuthHeader(s) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	token := s[1]

	if !isAuthenticated(token) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

func hasAuthHeader(s []string) bool {
	return len(s) != 2
}

func isAuthenticated(token string) bool {
    statusCode, _ := getUser(token)

    return statusCode == 200
}

func getUser(token string) (int,User) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", os.Getenv("AUTH_API_ME"), nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

    var body User
	io, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	if err := json.Unmarshal(io, &body); err != nil {
		fmt.Println(err)
	}

	return res.StatusCode, body
}
