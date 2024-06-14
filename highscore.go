package main

import (
	"encoding/json"
	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
	"log"
	"net/http"
)

func Highscore(client *http.Client) {

	var textLobby *walk.TextEdit
	var lobbyButton *walk.PushButton

	highscoreTake := func() {

		statsResponse := AskHS(client)
		jsonS, err := json.Marshal(statsResponse)
		if err != nil {
			log.Fatal(err)
		}
		jsonStringLiteral := string(jsonS)
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
				Text:      "Show Highscore",
				OnClicked: highscoreTake,
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}
}
