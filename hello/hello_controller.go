package hello

import (
	"but3/go-score/base"
	"but3/go-score/services/send"
	"math"
	"net/http"
	"strconv"
)

var Controller = base.Controller[HelloData]{
	Repository: Repository(),
}

// fonction invoquée pour répondre à une requête
func HelloWorld(res http.ResponseWriter, _ *http.Request) {
	// préparation des données à envoyer
	data := HelloWorldData{
		Message: "hello",
	}
	send.Json(res, data)
}

func Square(res http.ResponseWriter, r *http.Request) {
	// préparation des données à (envoyer
	value, _ := strconv.Atoi(r.PathValue("value"))
	data := HelloWorldData{
		Message: strconv.Itoa(int(math.Pow(float64(value), 2))),
	}
	send.Json(res, data)
}
