package main

import (
	"encoding/json"
	gui "github.com/grupawp/warships-lightgui/v2"
	"log"
	"net/http"
)

func Ask(client *http.Client, req *http.Request, board *gui.Board) Result {

	resp, err := client.Do(req)
	if err != nil {
		Ask(client, req, board)

	}
	var result Result
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Println("Error:", err)
	}

	return result
}

func Ask2(client *http.Client, req *http.Request, board *gui.Board) DescResult {

	resp, err := client.Do(req)
	if err != nil {
		Ask2(client, req, board)

	}
	var descResult DescResult
	err = json.NewDecoder(resp.Body).Decode(&descResult)
	if err != nil {
		log.Println("Error:", err)

	}

	return descResult
}
