package proxy

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const duration = 20 * time.Second

func Proxy(c *gin.Context, serviceURL string, path string) error {
	ctx, cancel := context.WithTimeout(c.Request.Context(), duration)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, c.Request.Method, serviceURL+path+"?"+c.Request.URL.RawQuery, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Internal Error"})
		return err
	}

	req.Header = c.Request.Header
    
	req.Header.Set("Content-Type", c.Request.Header.Get("Content-Type"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if err == context.DeadlineExceeded {
			c.JSON(http.StatusGatewayTimeout, gin.H{"error": "Request timed out"})
		} else {
			c.JSON(http.StatusBadGateway, gin.H{"msg": "Bad Gateway"})
		}

		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"msg": "Bad Gateway"})
		return err
	}

	c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), io.NopCloser(bytes.NewBuffer(body)), nil)

	return nil
}
