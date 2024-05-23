package main

import (
	"fmt"
	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
	"log"
	"strconv"
	"strings"
)

func shipCords() {
	var cords []string = make([]string, 20)
	var textCords, textFirst4, textLast4, textFirst31, textLast31, textFirst32, textLast32, textFirst21, textLast21, textFirst22, textLast22, textFirst23, textLast23, text11, text12, text13, text14 *walk.TextEdit
	var cordsButton *walk.PushButton

	onClickAddCords := func() {

		bigLetter := string(textFirst4.Text()[0])
		bigNumber, err := strconv.Atoi(string(textFirst4.Text()[1]))
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return
		}

		textSecond31 := string(textFirst31.Text()[0])
		secondNumber31, err := strconv.Atoi(string(textFirst31.Text()[1]))
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return
		}

		textFirst32.Text()
		textSecond32 := string(textFirst32.Text()[0])
		secondNumber32, err := strconv.Atoi(string(textFirst32.Text()[1]))
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return
		}

		cords[0] = textFirst4.Text()
		cords[1] = bigLetter + strconv.Itoa(bigNumber+1)
		cords[2] = bigLetter + strconv.Itoa(bigNumber+2)
		cords[3] = textLast4.Text()

		cords[4] = textFirst31.Text()
		cords[5] = textSecond31 + strconv.Itoa(secondNumber31+1)
		cords[6] = textLast31.Text()

		cords[7] = textFirst32.Text()
		cords[8] = textSecond32 + strconv.Itoa(secondNumber32+1)
		cords[9] = textLast32.Text()

		cords[10] = textFirst21.Text()
		cords[11] = textLast21.Text()

		cords[12] = textFirst22.Text()
		cords[13] = textLast22.Text()

		cords[14] = textFirst23.Text()
		cords[15] = textLast23.Text()

		cords[16] = text11.Text()
		cords[17] = text12.Text()
		cords[18] = text13.Text()
		cords[19] = text14.Text()

		err2 := textCords.SetText(strings.Join(cords, ", "))
		if err2 != nil {
			fmt.Println("Error setting text:", err2)
			return
		}
	}

	if _, err := (declarative.MainWindow{
		Title:  "Koordynaty statków",
		Size:   declarative.Size{Width: 450, Height: 300},
		Layout: declarative.VBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:     "cztero czesciowy ",
				ReadOnly: true,
			},
			declarative.Composite{
				Layout: declarative.HBox{},
				Children: []declarative.Widget{

					declarative.TextEdit{
						AssignTo: &textFirst4,
						ReadOnly: false,
					},
					declarative.TextEdit{
						AssignTo: &textLast4,
						ReadOnly: false,
					},
				},
			}, declarative.TextEdit{
				Text:     "pierwszy trzy czesciowy ",
				ReadOnly: true,
			},
			declarative.Composite{
				Layout: declarative.HBox{},
				Children: []declarative.Widget{

					declarative.TextEdit{
						AssignTo: &textFirst31,
						ReadOnly: false,
					},
					declarative.TextEdit{
						AssignTo: &textLast31,
						ReadOnly: false,
					},
				},
			}, declarative.TextEdit{
				Text:     "drugi trzy czesciowy ",
				ReadOnly: true,
			},
			declarative.Composite{
				Layout: declarative.HBox{},
				Children: []declarative.Widget{

					declarative.TextEdit{
						AssignTo: &textFirst32,
						ReadOnly: false,
					},
					declarative.TextEdit{
						AssignTo: &textLast32,
						ReadOnly: false,
					},
				},
			}, declarative.TextEdit{
				Text:     "pierwszy dwu czesciowy ",
				ReadOnly: true,
			},
			declarative.Composite{
				Layout: declarative.HBox{},
				Children: []declarative.Widget{

					declarative.TextEdit{
						AssignTo: &textFirst21,
						ReadOnly: false,
					},
					declarative.TextEdit{
						AssignTo: &textLast21,
						ReadOnly: false,
					},
				},
			}, declarative.TextEdit{
				Text:     "drugi dwu czesciowy ",
				ReadOnly: true,
			},
			declarative.Composite{
				Layout: declarative.HBox{},
				Children: []declarative.Widget{

					declarative.TextEdit{
						AssignTo: &textFirst22,
						ReadOnly: false,
					},
					declarative.TextEdit{
						AssignTo: &textLast22,
						ReadOnly: false,
					},
				},
			},
			declarative.TextEdit{
				Text:     "trzeci dwu czesciowy ",
				ReadOnly: true,
			},
			declarative.Composite{
				Layout: declarative.HBox{},
				Children: []declarative.Widget{

					declarative.TextEdit{
						AssignTo: &textFirst23,
						ReadOnly: false,
					},
					declarative.TextEdit{
						AssignTo: &textLast23,
						ReadOnly: false,
					},
				},
			},
			declarative.Composite{
				Layout: declarative.HBox{},
				Children: []declarative.Widget{
					declarative.TextEdit{
						Text:     "pierwsza jedynka ",
						ReadOnly: true,
					},
					declarative.TextEdit{
						AssignTo: &text11,
						ReadOnly: false,
					},
				},
			},
			declarative.Composite{
				Layout: declarative.HBox{},
				Children: []declarative.Widget{
					declarative.TextEdit{
						Text:     "druga jedynka ",
						ReadOnly: true,
					},
					declarative.TextEdit{
						AssignTo: &text12,
						ReadOnly: false,
					},
				},
			},
			declarative.Composite{
				Layout: declarative.HBox{},
				Children: []declarative.Widget{
					declarative.TextEdit{
						Text:     "trzecia jedynka ",
						ReadOnly: true,
					},
					declarative.TextEdit{
						AssignTo: &text13,
						ReadOnly: false,
					},
				},
			},
			declarative.Composite{
				Layout: declarative.HBox{},
				Children: []declarative.Widget{
					declarative.TextEdit{
						Text:     "czwarta jedynka ",
						ReadOnly: true,
					},
					declarative.TextEdit{
						AssignTo: &text14,
						ReadOnly: false,
					},
				},
			},
			declarative.PushButton{
				AssignTo:  &cordsButton,
				Text:      "Dodaj cordy",
				OnClicked: onClickAddCords,
			},
			declarative.TextEdit{
				AssignTo: &textCords,
				ReadOnly: true,
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}

}
