package main

import "github.com/lxn/walk"

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

type DescResult struct {
	Desc     string `json:"desc"`
	Nick     string `json:"nick"`
	OppDesc  string `json:"opp_desc"`
	Opponent string `json:"opponent"`
}

type CustomWidget struct {
	*walk.CustomWidget
	colors map[string]walk.Color
}
