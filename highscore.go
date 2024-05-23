package main

import (
	"encoding/json"
	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
	"log"
	"net/http"
)

func highscore(client *http.Client) {
	var textLobby *walk.TextEdit
	var lobbyButton *walk.PushButton
	highscoreTake := func() {
		url := "https://go-pjatk-server.fly.dev/api/stats"
		r, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println("Error:", err)
		}
		resp, err := client.Do(r)
		if err != nil {
			panic(err)
		}

		var statsResponse StatsResponse
		err = json.NewDecoder(resp.Body).Decode(&statsResponse)
		if err != nil {
			log.Println("Error:", err)
		}
		jsonString, err := json.Marshal(statsResponse)
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
		Title:  "Highscore",
		Size:   declarative.Size{Width: 450, Height: 300},
		Layout: declarative.VBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				AssignTo: &textLobby,
				ReadOnly: true,
			},
			declarative.PushButton{
				AssignTo:  &lobbyButton,
				Text:      "Show highscore",
				OnClicked: highscoreTake,
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}
}
