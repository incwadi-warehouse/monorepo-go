package web

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalln(err)
	}

	responseBody := bytes.NewBuffer(jsonData)

	client := http.Client{}
	req, err := http.NewRequest("POST", os.Getenv("MEILI")+"/indexes/books/search", responseBody)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("MEILI_TOKEN"))

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var v any
	err2 := json.Unmarshal(body, &v)
	if err2 != nil {
		log.Fatalln(err2)
	}

	c.JSON(res.StatusCode, v)
}

func List(c *gin.Context) {
	client := http.Client{}
	req, err := http.NewRequest("GET", os.Getenv("MEILI")+"/indexes", nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("MEILI_TOKEN"))

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var v any
	err2 := json.Unmarshal(body, &v)
	if err2 != nil {
		log.Fatalln(err2)
	}

	c.JSON(res.StatusCode, v)
}

func Create(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalln(err)
	}

	responseBody := bytes.NewBuffer(jsonData)

	client := http.Client{}
	req, err := http.NewRequest("POST", os.Getenv("MEILI")+"/indexes", responseBody)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("MEILI_TOKEN"))

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var v any
	err2 := json.Unmarshal(body, &v)
	if err2 != nil {
		log.Fatalln(err2)
	}

	c.JSON(res.StatusCode, v)
}

func Remove(c *gin.Context) {
	client := http.Client{}
	req, err := http.NewRequest("DELETE", os.Getenv("MEILI")+"/indexes/"+c.Param("index"), nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("MEILI_TOKEN"))

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var v any
	err2 := json.Unmarshal(body, &v)
	if err2 != nil {
		log.Fatalln(err2)
	}

	c.JSON(res.StatusCode, v)
}

func RemoveDocuments(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalln(err)
	}

	responseBody := bytes.NewBuffer(jsonData)

	client := http.Client{}
	req, err := http.NewRequest("DELETE", os.Getenv("MEILI")+"/indexes/books/documents", responseBody)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("MEILI_TOKEN"))

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var v any
	err2 := json.Unmarshal(body, &v)
	if err2 != nil {
		log.Fatalln(err2)
	}

	c.JSON(res.StatusCode, v)
}

func CreateDocument(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalln(err)
	}

	responseBody := bytes.NewBuffer(jsonData)

	client := http.Client{}
	req, err := http.NewRequest("POST", os.Getenv("MEILI")+"/indexes/books/documents", responseBody)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("MEILI_TOKEN"))

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var v any
	err2 := json.Unmarshal(body, &v)
	if err2 != nil {
		log.Fatalln(err2)
	}

	c.JSON(res.StatusCode, v)
}

func GetSettings(c *gin.Context) {
	client := http.Client{}
	req, err := http.NewRequest("GET", os.Getenv("MEILI")+"/indexes/books/settings", nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("MEILI_TOKEN"))

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var v any
	err2 := json.Unmarshal(body, &v)
	if err2 != nil {
		log.Fatalln(err2)
	}

	c.JSON(res.StatusCode, v)
}

func UpdateSettings(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalln(err)
	}

	responseBody := bytes.NewBuffer(jsonData)

	client := http.Client{}
	req, err := http.NewRequest("PATCH", os.Getenv("MEILI")+"/indexes/books/settings", responseBody)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("MEILI_TOKEN"))

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var v any
	err2 := json.Unmarshal(body, &v)
	if err2 != nil {
		log.Fatalln(err2)
	}

	c.JSON(res.StatusCode, v)
}


