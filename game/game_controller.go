package game

import (
	"but3/go-score/base"
	"but3/go-score/player"
	"but3/go-score/services/send"
	"encoding/json"
	"fmt"
	"net/http"
)

var Controller = base.Controller[Game]{
	Repository: Repository(),
}

type NewGameStruct struct {
	Players struct {
		Id1 string `json:"id1"`
		Id2 string `json:"id2"`
	} `json:"players"`
	Config struct {
		Tour player.Tour `json:"tour"`
		Sets int         `json:"sets"`
	} `json:"config"`
}

func NewGame(res http.ResponseWriter, r *http.Request) {
	fmt.Println("coucou")
	data := NewGameStruct{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return
	}
	player1, err := player.Repository().FindById(data.Players.Id1)
	if err != nil {
		return
	}
	player2, err := player.Repository().FindById(data.Players.Id2)
	if err != nil {
		return
	}
	game := Game{
		Players: Players{
			Player1: player1,
			Player2: player2,
		},
		Config: Config{
			Tour: data.Config.Tour,
			Sets: data.Config.Sets,
		},
	}
	err = Repository().InsertOne(game)
	if err != nil {
		return
	}
	send.Json(res, game)
}
