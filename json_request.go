package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

var requestClient = http.Client{Timeout: 10 * time.Second}

func GetJson(url string) (map[string]interface{}, error) {
	response, err := requestClient.Get(url)
	if err != nil {
		log.Println("Failed to send request\n", err)
		return nil, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Failed to read response\n", err)
		return nil, err
	}

	var jsonResponse map[string]interface{}
	err = json.Unmarshal(data, &jsonResponse)
	if err != nil {
		log.Println("Failed to parse response\n", err)
		return nil, err
	}

	return jsonResponse, nil
}

func GetJsonInStruct(url string, target interface{}) error {
	response, err := requestClient.Get(url)
	if err != nil {
		log.Println("Failed to send request\n", err)
		return err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Failed to read response\n", err)
		return err
	}

	//log.Println(string(data))

	err = json.Unmarshal(data, target)
	if err != nil {
		log.Println("Faield to parse response\n", err)
		return err
	}

	return nil
}
