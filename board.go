package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func Board(client *http.Client, dataRaw string) []string {

	url := "https://go-pjatk-server.fly.dev/api/game/board"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error:", err)
	}
	req.Header.Add("X-Auth-Token", dataRaw)
	resp, err := client.Do(req)
	if err != nil {
		Board(client, dataRaw)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	var result BoardResult

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Println("Error:", err)

	}

	data := result.Board
	return data

}
