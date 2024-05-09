package main

import (
	"bytes"
	"log"
	"net/http"
)

func initGame(client *http.Client, bodyText []byte) string {
	log.Println("init initgame")
	url := "https://go-pjatk-server.fly.dev/api/game"

	r, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyText))
	if err != nil {
		log.Println("Error:", err)
	}
	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	data := resp.Header.Get("X-Auth-Token")
	return data
}
