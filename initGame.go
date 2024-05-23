package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func initGame(client *http.Client, bodyText BodyText) string {
	log.Println("init initgame")
	url := "https://go-pjatk-server.fly.dev/api/game"

	b, err := json.Marshal(bodyText)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(string(b))
	r, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(fmt.Sprintf("%v", string(b)))))
	if err != nil {
		log.Println("Error:", err)
	}
	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	data := resp.Header.Get("X-Auth-Token")
	fmt.Println(data)
	return data
}
