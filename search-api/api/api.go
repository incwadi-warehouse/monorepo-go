package api

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func NewRequest(method, path string, requestBody io.Reader) (int, any) {
	body, err := getBody(requestBody)
	if err != nil {
		log.Fatalln(err)
	}

	res, err := doRequest(method, path, body)
	if err != nil {
		log.Fatalln(err)
	}

	data, err := getData(res)
	if err != nil {
		log.Fatalln(err)
	}

	return res.StatusCode, data
}

func NewRequestWithPlainRes(method, path string, requestBody io.Reader) *http.Response {
	body, err := getBody(requestBody)
	if err != nil {
		log.Fatalln(err)
	}

	res, err := doRequest(method, path, body)
	if err != nil {
		log.Fatalln(err)
	}

	return res
}

func setHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("MEILI_TOKEN"))
}

func getBody(requestBody io.Reader) (*bytes.Buffer, error) {
	jsonData, err := io.ReadAll(requestBody)
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(jsonData)

	return body, nil
}

func getData(res *http.Response) (any, error) {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data any
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func doRequest(method, path string, body *bytes.Buffer) (*http.Response, error) {
	client := http.Client{}
	req, err := http.NewRequest(method, os.Getenv("MEILI")+path, body)
	if err != nil {
		return nil, err
	}

	setHeaders(req)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
