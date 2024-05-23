package main

import (
	"encoding/json"
	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
	"log"
	"net/http"
	"time"
)

func lobby(client *http.Client) {
	var textLobby *walk.TextEdit
	var lobbyButton, highscoreButton *walk.PushButton
	onClickHighscore := func() {
		highscore(client)
	}
	onClickLobby := func() {
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

		err3 := textLobby.SetText(jsonStringLiteral)
		if err3 != nil {
			return
		}
	}
	if _, err := (declarative.MainWindow{
		Title:  "Lobby",
		Size:   declarative.Size{Width: 450, Height: 300},
		Layout: declarative.VBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				AssignTo: &textLobby,
				ReadOnly: true,
			},
			declarative.PushButton{
				AssignTo:  &highscoreButton,
				Text:      "Show Highscore",
				OnClicked: onClickHighscore,
			},
			declarative.PushButton{
				AssignTo:  &lobbyButton,
				Text:      "Show Lobby",
				OnClicked: onClickLobby,
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}
	go func() {
		for x := 1; x <= 360; x++ {
			url := "https://go-pjatk-server.fly.dev/api/lobby"
			r, err := http.NewRequest("POST", url, nil)
			if err != nil {
				log.Println("Error:", err)
			}
			resp, err := client.Do(r)
			if err != nil {
				panic(err)
			}
			var jsonData map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&jsonData)
			if err != nil {
				log.Println("Error:", err)
			}
			jsonString, err := json.Marshal(jsonData)
			if err != nil {
				log.Fatal(err)
			}

			jsonStringLiteral := string(jsonString)

			err3 := textLobby.SetText(jsonStringLiteral)
			if err3 != nil {
				return
			}
			time.Sleep(time.Second)
		}
	}()
}
