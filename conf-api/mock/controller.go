package mock

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
    Status int `json:"status"`
	Message string `json:"message"`
}

func Show(c *gin.Context) {
	if c.GetHeader("Authorization") == "Bearer token" {
		c.JSON(200, Response{200, c.GetHeader("Authorization")})
		return
	}

	c.JSON(401, Response{401, "UNAUTHORIZED"})
}
