package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func initGame(client *http.Client, bodyText BodyText) string {
	url := "https://go-pjatk-server.fly.dev/api/game"

	b, err := json.Marshal(bodyText)
	if err != nil {
		log.Println(err)
		return ""
	}
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {

		log.Println("Error:", err)
	}
	resp, err := client.Do(r)
	if err != nil {
		initGame(client, bodyText)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	data := resp.Header.Get("X-Auth-Token")
	return data
}
