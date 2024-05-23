package main

import (
	"github.com/fatih/color"
	gui "github.com/grupawp/warships-lightgui/v2"
	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Stats struct {
	Games  int    `json:"games"`
	Nick   string `json:"nick"`
	Points int    `json:"points"`
	Rank   int    `json:"rank"`
	Wins   int    `json:"wins"`
}

type BodyText struct {
	Coords     []string `json:"coords"`
	Desc       string   `json:"desc"`
	Nick       string   `json:"nick"`
	TargetNick string   `json:"target_nick"`
	WpBot      bool     `json:"wpbot"`
}

type StatsResponse struct {
	Stats []Stats `json:"stats"`
}
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

func main() {
	var textPlayer, textTimer, textEnemy, textStatus, fireLocation, textDesc, textView4 *walk.TextEdit
	var fireButton, leaveButton, startButton, lobbyWindowButton, cordsButton *walk.PushButton
	var checkBox *walk.CheckBox
	var fireText string
	client := &http.Client{}
	bodyText := BodyText{
		Coords:     make([]string, 20),
		Desc:       "",
		Nick:       "",
		TargetNick: "",
		WpBot:      true,
	}

	data := ""
	cfg := gui.NewConfig()
	cfg.HitChar = '#'
	cfg.MissChar = '&'
	cfg.MissColor = color.BgCyan
	cfg.HitColor = color.FgRed
	cfg.BorderColor = color.BgRed
	cfg.RulerTextColor = color.BgYellow
	board := gui.New(cfg)

	onClickLeave := func() {
		url := "https://go-pjatk-server.fly.dev/api/game/abandon"
		r, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			log.Println("Error:", err)
		}
		r.Header.Add("X-Auth-Token", data)
		_, err = client.Do(r)
		if err != nil {
			panic(err)
		}
	}
	onClickStart := func() {

		bodyText.Desc = textDesc.Text()
		bodyText.TargetNick = textEnemy.Text()
		bodyText.Nick = textPlayer.Text()

		err := textPlayer.SetReadOnly(true)
		if err != nil {
			return
		}
		err = textEnemy.SetReadOnly(true)
		if err != nil {
			return
		}
		err = textDesc.SetReadOnly(true)
		if err != nil {
			return
		}
		data = initGame(client, bodyText)
		startButton.SetVisible(false)
		checkBox.SetVisible(false)
		lobbyWindowButton.SetVisible(false)
		leaveButton.SetVisible(true)

		coords := Board(client, data)
		err2 := board.Import(coords)
		if err2 != nil {
		}
		url := "https://go-pjatk-server.fly.dev/api/game"
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println("Error:", err)
		}
		req.Header.Add("X-Auth-Token", data)
		go func() {
			for x := 1; x <= 360; x++ {
				responseText := Result{
					Status:     "",
					Nick:       "",
					LGS:        "",
					OppShots:   nil,
					Opponent:   "",
					ShouldFire: false,
					Timer:      0,
				}
				responseText = Ask(client, req, board)
				err := textTimer.SetText(strconv.Itoa(responseText.Timer))
				if err != nil {
					log.Println("Error setting text:", err)
				}

				err = textEnemy.SetText(responseText.Opponent)
				if err != nil {
					log.Println("Error setting text:", err)
				}
				err = textStatus.SetText(responseText.Status)
				if err != nil {
					log.Println("Error setting text:", err)
				}
				err = textPlayer.SetText(responseText.Nick)
				if err != nil {
					log.Println("Error setting text:", err)
				}
				time.Sleep(time.Second)
			}
		}()
		board.Display()
	}
	onClickHuman := func() {
		readOnly := textEnemy.ReadOnly()
		status := bodyText.WpBot
		if status {
			bodyText.WpBot = !status
		}
		err := textEnemy.SetReadOnly(!readOnly)
		if err != nil {
			return
		}
	}
	onClickFire := func() {
		fireText = Fire(client, fireLocation.Text(), data, board)
		err := textView4.SetText(fireText)
		if err != nil {
			return
		}
		board.Display()

	}

	onClickShowLobby := func() {
		lobby(client)
	}
	onClickCords := func() {
		shipCords(bodyText)
	}

	if _, err := (declarative.MainWindow{
		Title:  "Statki",
		Size:   declarative.Size{Width: 450, Height: 300},
		Layout: declarative.VBox{},
		Children: []declarative.Widget{
			declarative.PushButton{
				AssignTo:  &cordsButton,
				Text:      "Cords",
				OnClicked: onClickCords,
			},
			declarative.Composite{
				Layout: declarative.HBox{},
				Children: []declarative.Widget{
					declarative.TextEdit{
						AssignTo: &textPlayer,
						ReadOnly: false,
					},
					declarative.TextEdit{
						AssignTo: &textTimer,
						ReadOnly: true,
					},
					declarative.TextEdit{
						AssignTo: &textEnemy,
						ReadOnly: true,
					},
				},
			},
			declarative.CheckBox{
				AssignTo:  &checkBox,
				Text:      "Gra z człowiekiem?",
				OnClicked: onClickHuman,
			},
			declarative.TextEdit{
				AssignTo: &textDesc,
				ReadOnly: false,
			},
			declarative.TextEdit{
				AssignTo: &textStatus,
				ReadOnly: true,
			},
			declarative.PushButton{
				AssignTo:  &startButton,
				Text:      "Start",
				OnClicked: onClickStart,
			},

			declarative.PushButton{
				AssignTo:  &fireButton,
				Text:      "Fire",
				OnClicked: onClickFire,
			},
			declarative.PushButton{
				AssignTo:  &lobbyWindowButton,
				Text:      "ShowLobby",
				OnClicked: onClickShowLobby,
			},
			declarative.TextEdit{
				AssignTo: &fireLocation,
				ReadOnly: false,
			},
			declarative.TextEdit{
				AssignTo: &textView4,
				ReadOnly: true,
			},
			declarative.PushButton{
				AssignTo:  &leaveButton,
				Text:      "Leave",
				OnClicked: onClickLeave,
				Visible:   false,
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}

}
