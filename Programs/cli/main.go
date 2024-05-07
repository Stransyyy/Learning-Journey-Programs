package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/fatih/color"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Region  string `json:"region"`
		Country string `json:"country"`
		Time    string `json:"localtime"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		TempF     float64 `json:"temp_f"`
		Humidity  int     `json:"humidity"`
		UV        float64 `json:"uv"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
}

func main() {
	q := "Tupelo"

	if len(os.Args) >= 2 {
		q = os.Args[1]
	}

	res, err := http.Get("http://api.weatherapi.com/v1/current.json?key=e60f2ddf09bc4b3091c140905240705&q=" + q + "&aqi=no")
	if err != nil {
		fmt.Println("HTTP request failed: ", err)
		return
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			panic(err)
		}
	}()

	if res.StatusCode != http.StatusOK {
		fmt.Println("HTTP request failed: ", res.Status)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to read response body: ", err)
		return
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		fmt.Println("Failed to unmarshal JSON: ", err)
		return
	}

	location, current := weather.Location, weather.Current

	// Color settings for temperature
	tempCColor := color.New(color.FgCyan) // Default color
	if current.TempC >= 30 {
		tempCColor = color.New(color.FgRed)
	} else if current.TempC <= 0 {
		tempCColor = color.New(color.FgBlue)
	}
	tempCString := tempCColor.Sprintf("%.0fC°", current.TempC)

	tempFColor := color.New(color.FgCyan) // Default color
	if current.TempF >= 86 {
		tempFColor = color.New(color.FgRed)
	} else if current.TempF <= 32 {
		tempFColor = color.New(color.FgBlue)
	}
	tempFString := tempFColor.Sprintf("%.0fF°", current.TempF)

	// Color settings for humidity
	humidityColor := color.New(color.FgYellow) // Default color
	if current.Humidity >= 80 {
		humidityColor = color.New(color.FgRed)
	} else if current.Humidity <= 20 {
		humidityColor = color.New(color.FgBlue)
	}
	humidityString := humidityColor.Sprintf("%v%%", current.Humidity)

	// Color settings for UV
	uvColor := color.New(color.FgHiWhite)
	switch {
	case current.UV < 3:
		uvColor = color.New(color.FgGreen)
	case current.UV < 6:
		uvColor = color.New(color.FgYellow)
	case current.UV >= 6 && current.UV < 8:
		uvColor = color.New(color.FgRed)
	case current.UV >= 8 && current.UV < 11:
		uvColor = color.New(color.FgMagenta)
	case current.UV >= 11:
		uvColor = color.New(color.FgHiBlue)
	}

	uvStr := uvColor.Sprintf("%.0f", current.UV)

	// Constructing the message with color formatting
	message := fmt.Sprintf("Location: %s, %s, %s\nWeather: %s, %s\nCondition of the day: %s\nExtra: Humidity: %s, UV:",
		location.Name, location.Region, location.Country,
		tempCString, tempFString,
		current.Condition.Text,
		humidityString)

	fmt.Print(message + uvStr + " Time: " + location.Time + "\n")
}

func prettyJson(body []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, body, "", "  ")
	return out.Bytes(), err
}
