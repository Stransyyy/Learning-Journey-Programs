package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type File struct {
	User            string   `json:"user"`
	Password        string   `json:"password"`
	Server          string   `json:"server"`
	Port            int      `json:"port"`
	Protocol        string   `json:"protocol"`
	FilesToUpload   []string `json:"filesToUpload"`
	FilesToDownload []string `json:"filesToDownload"`
}

func jsonFileReader() {
	data, derr := os.ReadFile("file.json")

	if derr != nil {
		log.Fatal(derr)
	}

	var f File

	err := json.Unmarshal(data, &f)

	if err != nil {
		fmt.Println("error:", err)
	}

}

func main() {
	// load configuration file

	// connect to server

	// upload or download files
}
