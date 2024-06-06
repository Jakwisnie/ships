package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func initGame(client *http.Client, bodyText BodyText) string {
	log.Println("init initgame")
	url := "https://go-pjatk-server.fly.dev/api/game"

	b, err := json.Marshal(bodyText)
	if err != nil {
		log.Println(err)
		return ""
	}
	log.Println(string(b))
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {

		log.Println("Error:", err)
	}
	resp, err := client.Do(r)
	if err != nil {
		initGame(client, bodyText)
	}

	data := resp.Header.Get("X-Auth-Token")
	log.Println(resp)
	log.Println(data)
	return data
}
