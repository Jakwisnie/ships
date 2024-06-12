package main

import (
	"fmt"
	"github.com/fatih/color"
	gui "github.com/grupawp/warships-lightgui/v2"
	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func accuracyToString(accuracy float64) string {
	return fmt.Sprintf("Accuracy: %.2f%%", accuracy*100)
}
func main() {
	var mainWindow *walk.MainWindow
	var customWidget, customWidget2 *CustomWidget
	var textPlayer, textTimer, textEnemy, textEnemyDesc, textStatus, fireLocation, textDesc, textView4, textEnemyShots, textAccuracy *walk.TextEdit
	var fireButton, leaveButton, startButton, restartButton, lobbyWindowButton, cordsButton *walk.PushButton
	var checkBox *walk.CheckBox
	var fireText string
	var shotCount = 0
	var goodShot = 0
	var red = walk.RGB(255, 0, 0)
	var green = walk.RGB(0, 255, 0)
	stop := false
	client := &http.Client{}
	bodyText := BodyText{
		Coords:     make([]string, 20),
		Desc:       "",
		Nick:       "",
		TargetNick: "",
		WpBot:      true,
	}
	customWidget = &CustomWidget{
		colors: make(map[string]walk.Color),
	}
	customWidget2 = &CustomWidget{
		colors: make(map[string]walk.Color),
	}

	data := ""
	cfg := gui.NewConfig()
	cfg.HitChar = '#'
	cfg.MissChar = '&'
	cfg.MissColor = color.BgCyan
	cfg.HitColor = color.BgGreen
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
		stop = true
		time.Sleep(time.Second / 2)
		err2 := mainWindow.Invalidate()
		if err2 != nil {
			return
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
		err = textEnemyDesc.SetReadOnly(true)
		if err != nil {
			return
		}
		data = initGame(client, bodyText)
		startButton.SetVisible(false)
		checkBox.SetVisible(false)
		lobbyWindowButton.SetVisible(false)
		leaveButton.SetVisible(true)
		stop = false
		coords := Board(client, data)
		err2 := board.Import(coords)
		if err2 != nil {
		}
		url := "https://go-pjatk-server.fly.dev/api/game/desc"
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println("Error:", err)
		}
		req.Header.Add("X-Auth-Token", data)
		url2 := "https://go-pjatk-server.fly.dev/api/game"
		req2, err := http.NewRequest("GET", url2, nil)
		if err != nil {
			log.Println("Error:", err)
		}
		req2.Header.Add("X-Auth-Token", data)
		responseText := Result{
			Status:     "",
			Nick:       "",
			LGS:        "",
			OppShots:   nil,
			Opponent:   "",
			ShouldFire: false,
			Timer:      0,
		}
		descResult := DescResult{
			Desc:     "",
			Nick:     "",
			OppDesc:  "",
			Opponent: "",
		}
		go func() {
			for {
				time.Sleep(time.Second / 2)
				responseText = Ask(client, req2, board)
				var y = responseText.Timer
				if y != 0 {
					err := textTimer.SetText(strconv.Itoa(responseText.Timer))
					if err != nil {
						log.Println("Error setting text:", err)
					}
				}
				var z = responseText.Opponent
				if z != "" {
					err = textEnemy.SetText(responseText.Opponent)
					if err != nil {
						log.Println("Error setting text:", err)
					}
				}
				z = ""
				z = responseText.Status
				if z != "" {
					err = textStatus.SetText(responseText.Status)
					if err != nil {
						log.Println("Error setting text:", err)
					}
				}
				z = ""
				z = responseText.Nick
				if z != "" {
					err = textPlayer.SetText(responseText.Nick)
					if err != nil {
						log.Println("Error setting text:", err)
					}
				}
				z = ""
				z = responseText.Opponent
				if z != "" {
					err = textEnemy.SetText(responseText.Opponent)
					if err != nil {
						log.Println("Error setting text:", err)
					}
				}
				z = ""
				z = strings.Join(responseText.OppShots, ",")
				if z != "" {
					err = textEnemyShots.SetText(strings.Join(responseText.OppShots, ","))
					if err != nil {
						log.Println("Error setting text:", err)
					}
				}

				descResult = Ask2(client, req, board)
				time.Sleep(time.Second / 2)
				z = descResult.OppDesc
				if z != "" {
					err = textEnemyDesc.SetText(descResult.OppDesc)
					if err != nil {
						log.Println("Error setting text:", err)
					}
				}

				if stop == true {
					break
				}
			}
		}()

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
		fireText = Fire(client, strings.ToUpper(fireLocation.Text()), data, board)
		err := textView4.SetText(fireText)

		if err != nil {
			return
		}
		shotCount = shotCount + 1

		if fireText == "hit" || fireText == "sunk" {
			goodShot = goodShot + 1
			customWidget2.colors[strings.ToUpper(fireLocation.Text())] = green
		} else {
			customWidget2.colors[strings.ToUpper(fireLocation.Text())] = red
		}
		accuracy := float64(goodShot) / float64(shotCount)
		accuracyStr := accuracyToString(accuracy)
		err2 := textAccuracy.SetText(accuracyStr)
		if err2 != nil {
			return
		}

		arr := strings.Split(textEnemyShots.Text(), ",")
		for _, cord := range arr {
			customWidget.colors[cord] = walk.RGB(255, 0, 0)
		}
		err4 := customWidget.Invalidate()
		if err4 != nil {
			return
		}
		err3 := customWidget2.Invalidate()
		if err3 != nil {
			return
		}
	}
	onClickRestart := func() {
		customWidget.colors = make(map[string]walk.Color)
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
		checkBox.SetChecked(false)
		lobbyWindowButton.SetVisible(true)
		err = textPlayer.SetReadOnly(false)
		if err != nil {
			return
		}
		err = textEnemy.SetReadOnly(true)
		if err != nil {
			return
		}
		err = textDesc.SetReadOnly(false)
		if err != nil {
			return
		}
		err = textEnemyDesc.SetReadOnly(true)
		if err != nil {
			return
		}
		time.Sleep(time.Second / 2)
		err2 := mainWindow.Invalidate()
		if err2 != nil {
			return
		}
	}
	onClickShowLobby := func() {
		lobby(client)
	}
	onClickCords := func() {
		shipCords(bodyText, customWidget)
	}

	if _, err := (declarative.MainWindow{
		AssignTo: &mainWindow,
		Title:    "Statki",
		Size:     declarative.Size{Width: 850, Height: 850},
		Layout:   declarative.VBox{},

		Children: []declarative.Widget{
			//Top place
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
			//Remaings
			declarative.Composite{
				Layout: declarative.HBox{},

				StretchFactor: 4,
				Children: []declarative.Widget{
					//player text
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
							}, declarative.TextEdit{
								AssignTo: &textAccuracy,
								ReadOnly: true,
							},
							declarative.VSpacer{MinSize: declarative.Size{
								Width:  300,
								Height: 300,
							}},
						},
					},
					//mid text
					declarative.Composite{
						Layout:        declarative.VBox{},
						StretchFactor: 1,
						Border:        true,
						Children: []declarative.Widget{
							declarative.TextEdit{
								AssignTo:      &textTimer,
								ReadOnly:      true,
								StretchFactor: 1,
							},

							declarative.TextEdit{
								AssignTo:      &textStatus,
								ReadOnly:      true,
								StretchFactor: 1,
							},

							declarative.TextEdit{
								AssignTo:      &textView4,
								ReadOnly:      true,
								StretchFactor: 1},
							declarative.Composite{
								Layout:        declarative.HBox{},
								StretchFactor: 8,
								Border:        true,

								Children: []declarative.Widget{
									declarative.CustomWidget{
										AssignTo: &customWidget.CustomWidget,
										Paint:    customWidget.DrawFull,
									},
									declarative.CustomWidget{
										AssignTo: &customWidget2.CustomWidget,
										Paint:    customWidget2.DrawFull,
									},
								}},
							declarative.VSpacer{MinSize: declarative.Size{
								Width: 550,
							}},
						},
						//enemy text
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
							declarative.TextEdit{
								AssignTo: &textEnemyShots,
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
