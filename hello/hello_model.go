package hello

import (
	"but3/go-score/base"
)

// structure des données retournées par HelloWorld
type HelloWorldData struct {
	Message string
}

type HelloData struct {
	base.Entity `bson:"inline"`
	Message     string `bson:"message,string" json:message`
}
