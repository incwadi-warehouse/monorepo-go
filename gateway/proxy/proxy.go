package proxy

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Proxy forwards an HTTP request to a specified service URL and returns the response.
// It handles request creation, proxying, and response handling.
func Proxy(c *gin.Context, serviceURL string, path string) error {
	const duration = 20 * time.Second

	ctx, cancel := context.WithTimeout(c.Request.Context(), duration)
	defer cancel()

	req, err := request(ctx, c, serviceURL, path)
	if err != nil {
		return err
	}

	return response(req, c)
}

// request creates a new HTTP request based on the Gin context and target service details.
func request(ctx context.Context, c *gin.Context, serviceURL string, path string) (*http.Request, error) {
	fullURL := serviceURL + path + "?" + c.Request.URL.RawQuery

	req, err := http.NewRequestWithContext(ctx, c.Request.Method, fullURL, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Internal Error"})
		return nil, err
	}

	req.Header = make(http.Header)
	for key, values := range c.Request.Header {
		req.Header[key] = values
	}

	return req, nil
}

// response sends the created request to the target service and handles the response.
func response(req *http.Request, c *gin.Context) error {
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

	for key, value := range resp.Header {
		c.Header(key, value[0])
	}

	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)

	return nil
}
