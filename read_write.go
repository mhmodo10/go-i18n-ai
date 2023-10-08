package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// safely close file
func safeClose(f *os.File) {
	if err := f.Close(); err != nil {
		fmt.Println("failed to close file: %w", err)
	}
}

// read file and parse its contents to map[string]interface{}
func ReadJson(path string) (map[string]interface{}, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("could not read file: %w", err)
		return nil, err
	}
	defer safeClose(file)
	data := make(map[string]interface{})
	decodeErr := json.NewDecoder(file).Decode(&data)
	if decodeErr != nil {
		return nil, err
	}
	if data == nil {
		return make(map[string]interface{}), nil
	}
	return data, nil
}

// given a path and a map[string]interface{}
//
// marshal and write the map into the given path
func WriteResultToJson(path string, m map[string]interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer safeClose(file)

	jsonData, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		return err
	}

	// Write the JSON data to the file
	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}
	return nil
}
