package data

import (
	"encoding/json"
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

type Data interface {
	GetName() (string, error)

	GetFlag() (string, error)

	GetLanguages() ([]Languages, error)

	GetPopulation() (int, error)
}

func (d *Data) GetName() ([]Country, error) {
	response, err := http.Get("https://restcountries.eu/rest/v2/name")
	if err != nil {
		return nil, fmt.Errorf("Error fetching data from %s: %v", response.Request.URL, err)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response data: %v", err)
	}

	var data Country

	err = json.Unmarshal(responseData, &data)

}
