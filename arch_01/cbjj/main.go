package main

import (
	"encoding/json"
	"os"
)

type Data struct {
	// Define your JSON structure here
	// Example:
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

func readJSON(filename string) (Data, error) {
	var data Data

	jsonFile, err := os.Open(filename)
	if err != nil {
		return data, err
	}
	defer jsonFile.Close()

	byteValue, _ := os.ReadFile(jsonFile)

	err = json.Unmarshal(byteValue, &data)
	return data, err
}
