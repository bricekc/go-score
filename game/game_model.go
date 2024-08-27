package game

import (
	"but3/go-score/base"
	"but3/go-score/player"
)

type Players struct {
	Player1 player.Player `bson:"player1" json:"player1"`
	Player2 player.Player `bson:"player2" json:"player2"`
}

type Score struct {
	Sets   int   `bson:"score_sets,int" json:"score_sets"`
	Games  []int `bson:"games,int" json:"games"`
	Points int   `bson:"points,int" json:"points"`
}

type Config struct {
	Tour player.Tour `bson:"tour,string" json:"tour"`
	Sets int         `bson:"config_sets,int" json:"config_sets"`
}
type State struct {
	CurrentSet int     `bson:"currentSet,int" json:currentSet`
	TieBreak   bool    `bson:"tieBreak,bool" json:tieBreak`
	Scores     []Score `bson:"scores" json:scores`
	Winner     int     `bson:"winner,int" json:winner`
}

type Game struct {
	base.Entity `bson:"inline"`
	Players     Players `bson:"players" json:players`
	State       State   `bson:"state" json:state`
	Config      Config  `bson:"config" json:config`
}
