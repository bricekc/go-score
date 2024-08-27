package player

import (
	"but3/go-score/base"
	"but3/go-score/services/mongodb"
)

var repository = base.Repository[Player]{
	Collection: mongodb.Collection("player"),
}

func Repository() base.Repository[Player] {
	return repository
}
