package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	gui "github.com/grupawp/warships-lightgui/v2"
	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
	"io"
	"log"
	"net/http"
	"time"
)

type Result struct {
	Status     string   `json:"game_status"`
	Nick       string   `json:"nick"`
	LGS        string   `json:"last_game_status"`
	OppShots   []string `json:"opp_shots"`
	Opponent   string   `json:"opponent"`
	ShouldFire bool     `json:"should_fire"`
	Timer      int      `json:"timer"`
}
type ShootResult struct {
	Result string `json:"result"`
}
type BoardResult struct {
	Board []string `json:"board"`
}

func initGame(client *http.Client) string {
	log.Println("init initgame")
	url := "https://go-pjatk-server.fly.dev/api/game"
	var bodyText = []byte(`{
		 "coords": [
    "A1",
    "A3",
    "B9",
    "C7",
    "D1",
    "D2",
    "D3",
    "D4",
    "D7",
    "E7",
    "F1",
    "F2",
    "F3",
    "F5",
    "G5",
    "G8",
    "G9",
    "I4",
    "J4",
    "J8"
  ],
	"desc": "Ship game",
	"nick": "Abuk",
	"target_nick": "",
	"wpbot": true
	}`)
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyText))
	if err != nil {
		log.Println("Error:", err)
	}
	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	log.Println("Response from ", resp)
	data := resp.Header.Get("X-Auth-Token")
	log.Println("X-Auth-Token is ", data)
	return data
}
func Ask(client *http.Client, req *http.Request, board *gui.Board) string {

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error:", err)

	}
	var result Result
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Println("Error:", err)

	}

	response := fmt.Sprintf("Oponent is %s\n | ", result.Opponent)
	response += fmt.Sprintf("Nick is %s\n | ", result.Nick)
	response += fmt.Sprintf("Last Game Status is %s\n | ", result.LGS)
	response += fmt.Sprintf("Should Fire is %v\n | ", result.ShouldFire)
	response += fmt.Sprintf("Timer is %v\n | ", result.Timer)
	response += fmt.Sprintf("GameStatus is %s\n | ", result.Status)
	response += fmt.Sprintf("Oop Shots is %s\n | ", result.OppShots)
	if len(result.OppShots) > 0 {
		for x := 0; x < len(result.OppShots); x++ {
			err := board.Set(gui.Left, result.OppShots[x], gui.Miss)
			if err != nil {

			}
		}
	}
	return response
}
func Board(client *http.Client, data2 string) []string {

	url := "https://go-pjatk-server.fly.dev/api/game/board"
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error:", err)
	}
	r.Header.Add("X-Auth-Token", data2)
	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var result BoardResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Println("Error:", err)

	}

	data := result.Board
	return data

}
func Fire(client *http.Client, fireLocation string, data string, board *gui.Board) string {
	log.Println("Inite Fire")
	url := "https://go-pjatk-server.fly.dev/api/game/fire"

	var bodyText = []byte(`{"coord": "` + fireLocation + `"}`)
	log.Println(bytes.NewBuffer(bodyText))
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyText))
	r.Header.Add("X-Auth-Token", data)
	resp, err := client.Do(r)
	if err != nil {
		log.Println("Error:", err)

	}
	log.Println(resp)
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
func main() {
	var responseText, fireText string
	client := &http.Client{}
	data := initGame(client)
	log.Printf("Client on")
	cfg := gui.NewConfig()
	cfg.HitChar = '#'
	cfg.MissChar = '&'
	cfg.MissColor = color.BgCyan
	cfg.HitColor = color.FgRed
	cfg.BorderColor = color.BgRed
	cfg.RulerTextColor = color.BgYellow

	board := gui.New(cfg)
	coords := Board(client, data)
	err := board.Import(coords)
	if err != nil {

	}

	url := "https://go-pjatk-server.fly.dev/api/game"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error:", err)
	}
	req.Header.Add("X-Auth-Token", data)
	var textView, textView2, textView3, fireLocation *walk.TextEdit

	onClick2 := func() {

		board.Display()

	}
	onClick3 := func() {
		fireText = Fire(client, fireLocation.Text(), data, board)
		textView3.SetText(fireText)
		board.Display()
	}
	onClick4 := func() {
		go func() {
			for x := 1; x <= 360; x++ {
				responseText = Ask(client, req, board)
				err := textView.SetText(responseText)
				if err != nil {
					log.Println("Error setting text:", err)
				}

				time.Sleep(time.Second)
			}
		}()
	}

	if _, err := (declarative.MainWindow{
		Title:  "Statki",
		Size:   declarative.Size{Width: 450, Height: 300},
		Layout: declarative.VBox{},
		Children: []declarative.Widget{
			declarative.PushButton{
				Text:      "Ask for data",
				OnClicked: onClick4,
			},
			declarative.TextEdit{
				AssignTo: &textView,
				ReadOnly: true,
			},
			declarative.PushButton{
				Text:      "Ask for board",
				OnClicked: onClick2,
			},
			declarative.TextEdit{
				AssignTo: &textView2,
				ReadOnly: true,
			},
			declarative.PushButton{
				Text:      "Fire",
				OnClicked: onClick3,
			},
			declarative.TextEdit{
				AssignTo: &fireLocation,
				ReadOnly: false,
			},
			declarative.TextEdit{
				AssignTo: &textView3,
				ReadOnly: true,
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}

}
