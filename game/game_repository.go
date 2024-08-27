package game

import (
	"but3/go-score/base"
	"but3/go-score/services/mongodb"
)

var repository = base.Repository[Game]{
	Collection: mongodb.Collection("game"),
}

func Repository() base.Repository[Game] {
	return repository
}
