package main

import (
	"encoding/json"
	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
	"log"
	"net/http"
)

func lobby(client *http.Client) {
	var textLobby *walk.TextEdit
	var lobbyButton, highscoreButton *walk.PushButton
	onClickHighscore := func() {
		Highscore(client)
	}
	onClickLobby := func() {
		jsonData := AskLB(client)
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

}
