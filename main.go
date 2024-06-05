package main

import (
	"encoding/json"
	"github.com/fatih/color"
	gui "github.com/grupawp/warships-lightgui/v2"
	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	var textPlayer, textTimer, textEnemy, textEnemyDesc, textStatus, fireLocation, textDesc, textView4 *walk.TextEdit
	var fireButton, leaveButton, startButton, restartButton, lobbyWindowButton, cordsButton *walk.PushButton
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
		restartButton.SetVisible(true)
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
		err = textEnemyDesc.SetReadOnly(true)
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
		time.Sleep(time.Second * 5)
		url = "https://go-pjatk-server.fly.dev/api/game"
		req, err = http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println("Error:", err)
		}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error:", err)

		}
		var result DescResult
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			log.Println("Error:", err)

		}
		err = textEnemyDesc.SetText(result.OppDesc)
		if err != nil {
			log.Println("Error setting text:", err)
		}

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
	onClickRestart := func() {
		err := textPlayer.SetText("")
		if err != nil {
			return
		}
		err = textTimer.SetText("")
		if err != nil {
			return
		}
		err = textEnemy.SetText("")
		if err != nil {
			return
		}
		err = textEnemyDesc.SetText("")
		if err != nil {
			return
		}
		err = textStatus.SetText("")
		if err != nil {
			return
		}
		err = fireLocation.SetText("")
		if err != nil {
			return
		}
		err = textDesc.SetText("")
		if err != nil {
			return
		}
		err = textView4.SetText("")
		if err != nil {
			return
		}
		startButton.SetVisible(true)
		restartButton.SetVisible(false)
		leaveButton.SetVisible(false)
		checkBox.SetVisible(true)
		lobbyWindowButton.SetVisible(true)
		err = textPlayer.SetReadOnly(false)
		if err != nil {
			return
		}
		err = textEnemy.SetReadOnly(false)
		if err != nil {
			return
		}
		err = textDesc.SetReadOnly(false)
		if err != nil {
			return
		}
		err = textEnemyDesc.SetReadOnly(false)
		if err != nil {
			return
		}
	}
	onClickShowLobby := func() {
		lobby(client)
	}
	onClickCords := func() {
		shipCords(bodyText)
	}

	if _, err := (declarative.MainWindow{
		Title:  "Statki",
		Size:   declarative.Size{Width: 850, Height: 640},
		Layout: declarative.VBox{},

		Children: []declarative.Widget{

			declarative.Composite{
				Layout:        declarative.HBox{Alignment: declarative.AlignHCenterVCenter},
				StretchFactor: 1,
				Border:        true,
				Children: []declarative.Widget{
					declarative.Composite{
						Layout: declarative.Grid{Columns: 5},

						Children: []declarative.Widget{
							declarative.PushButton{
								AssignTo:  &startButton,
								Text:      "Start",
								OnClicked: onClickStart,
							},
							declarative.PushButton{
								AssignTo:  &restartButton,
								Text:      "Restart",
								OnClicked: onClickRestart,
								Visible:   false,
							},
							declarative.PushButton{
								AssignTo:  &leaveButton,
								Text:      "Leave",
								OnClicked: onClickLeave,
								Visible:   false,
							},
							declarative.PushButton{
								AssignTo:  &cordsButton,
								Text:      "Cords",
								OnClicked: onClickCords,
							},
							declarative.PushButton{
								AssignTo:  &lobbyWindowButton,
								Text:      "ShowLobby",
								OnClicked: onClickShowLobby,
							}}},
					declarative.Composite{
						Layout: declarative.Grid{Columns: 2},
						Border: true,
						Children: []declarative.Widget{
							declarative.TextEdit{
								AssignTo: &fireLocation,
								ReadOnly: false,
								MaxSize:  declarative.Size{Width: 150, Height: 25},
							},
							declarative.PushButton{
								AssignTo:  &fireButton,
								Text:      "Fire",
								OnClicked: onClickFire,
							}}},
				},
			},
			declarative.Composite{
				Layout:        declarative.HBox{},
				StretchFactor: 4,
				Children: []declarative.Widget{
					declarative.Composite{
						Layout:        declarative.VBox{},
						StretchFactor: 1,
						Border:        true,
						Children: []declarative.Widget{
							declarative.TextEdit{
								AssignTo: &textPlayer,
								ReadOnly: false,
							},
							declarative.TextEdit{
								AssignTo: &textDesc,
								ReadOnly: false,
							}, declarative.VSpacer{MinSize: declarative.Size{
								Width:  300,
								Height: 300,
							}},
						},
					},
					declarative.Composite{
						Layout:        declarative.VBox{},
						StretchFactor: 1,
						Border:        true,
						Children: []declarative.Widget{
							declarative.TextEdit{
								AssignTo: &textTimer,
								ReadOnly: true,
							},

							declarative.TextEdit{
								AssignTo: &textStatus,
								ReadOnly: true,
							},

							declarative.TextEdit{
								AssignTo: &textView4,
								ReadOnly: true}, declarative.VSpacer{MinSize: declarative.Size{
								Width:  300,
								Height: 300,
							}},
						},
					}, declarative.Composite{
						Layout:        declarative.VBox{},
						StretchFactor: 1,
						Border:        true,
						Children: []declarative.Widget{
							declarative.TextEdit{
								AssignTo: &textEnemy,
								ReadOnly: true,
							},
							declarative.TextEdit{
								AssignTo: &textEnemyDesc,
								ReadOnly: true,
							},

							declarative.CheckBox{
								AssignTo:      &checkBox,
								Text:          "Human?",
								OnClicked:     onClickHuman,
								StretchFactor: 1,
							},
							declarative.VSpacer{MinSize: declarative.Size{
								Width:  300,
								Height: 300,
							}},
						},
					},
				}},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}

}
