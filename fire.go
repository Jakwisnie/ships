package main

import (
	"bytes"
	"encoding/json"
	gui "github.com/grupawp/warships-lightgui/v2"
	"io"
	"log"
	"net/http"
)

func Fire(client *http.Client, fireLocation string, data string, board *gui.Board) string {
	url := "https://go-pjatk-server.fly.dev/api/game/fire"

	var bodyText = []byte(`{"coord": "` + fireLocation + `"}`)
	log.Println(bytes.NewBuffer(bodyText))
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyText))
	r.Header.Add("X-Auth-Token", data)
	resp, err := client.Do(r)
	if err != nil {
		Fire(client, fireLocation, data, board)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	var result ShootResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Println("Error:", err)

	}
	switch os := result.Result; os {
	case "hit":
		err := board.Set(gui.Right, fireLocation, gui.Hit)
		if err != nil {
			return ""
		}
	case "miss":
		err := board.Set(gui.Right, fireLocation, gui.Miss)
		if err != nil {
			return ""
		}
	case "ship":
		err := board.Set(gui.Right, fireLocation, gui.Ship)
		if err != nil {
			return ""
		}
	default:
	}

	return result.Result
}
