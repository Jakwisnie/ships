package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func Board(client *http.Client, data2 string) []string {

	url := "https://go-pjatk-server.fly.dev/api/game/board"
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error:", err)
	}
	r.Header.Add("X-Auth-Token", data2)
	resp, err := client.Do(r)
	if err != nil {
		panic(err)
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
