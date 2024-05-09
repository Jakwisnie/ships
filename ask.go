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
		log.Println("Error:", err)

	}
	var result Result
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Println("Error:", err)

	}

	return result
}
