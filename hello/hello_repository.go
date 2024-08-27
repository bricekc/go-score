package hello

import (
	"but3/go-score/base"
	"but3/go-score/services/mongodb"
)

var repository = base.Repository[HelloData]{
	Collection: mongodb.Collection("hello"),
}

func Repository() base.Repository[HelloData] {
	return repository
}
