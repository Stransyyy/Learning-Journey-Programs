package data

import (
	"fmt"
	"io"
	"net/http"
)

type Country struct {
	Name       string `json:"name"`
	Flag       string `json:"flag"`
	Languages  []Languages
	Population int `json:"population"`
}

type Languages struct {
	Name string `json:"name"`
}

func get() (string, error) {
	response, err := http.Get("https://restcountries.eu/rest/v2/all")
	if err != nil {
		return "", fmt.Errorf("Error fetching data from %s: %v", response.Request.URL, err)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading response data: %v", err)
	}

	return string(responseData), nil
}
