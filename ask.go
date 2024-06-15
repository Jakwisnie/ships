package main

import (
	"encoding/json"
	gui "github.com/grupawp/warships-lightgui/v2"
	"io"
	"log"
	"net/http"
)

func AskBase(client *http.Client, board *gui.Board, data string) Result {
	url := "https://go-pjatk-server.fly.dev/api/game"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error:", err)
	}
	req.Header.Add("X-Auth-Token", data)
	resp, err := client.Do(req)
	if err != nil {
		AskBase(client, board, data)

	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	var result Result
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Println("Error:", err)
	}

	return result
}

func AskDesc(client *http.Client, board *gui.Board, data string) DescResult {
	url := "https://go-pjatk-server.fly.dev/api/game/desc"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error:", err)
	}
	req.Header.Add("X-Auth-Token", data)
	resp, err := client.Do(req)
	if err != nil {
		AskDesc(client, board, data)

	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	var descResult DescResult
	err = json.NewDecoder(resp.Body).Decode(&descResult)
	if err != nil {
		log.Println("Error:", err)

	}

	return descResult
}
func AskLB(client *http.Client) string {
	url := "https://go-pjatk-server.fly.dev/api/lobby"
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error:", err)
	}
	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	var jsonData []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&jsonData)
	if err != nil {
		log.Println("Error:", err)
	}
	jsonString, err := json.Marshal(jsonData)
	if err != nil {
		log.Fatal(err)
	}

	jsonStringLiteral := string(jsonString)
	return jsonStringLiteral
}
func AskHS(client *http.Client) StatsResponse {
	url := "https://go-pjatk-server.fly.dev/api/stats"
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error:", err)
	}
	resp, err := client.Do(r)
	if err != nil {
		AskHS(client)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var statsResponse StatsResponse
	err = json.NewDecoder(resp.Body).Decode(&statsResponse)
	if err != nil {
		log.Println("Error:", err)
	}

	return statsResponse
}
