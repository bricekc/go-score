package base

import (
	"but3/go-score/services/send"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

// Définition d'un type représentant une fonction qui génère un filtre
// à partir des données présentes dans la requête http
type CreateFilterFunction func(*http.Request) bson.D

type Controller[Entity EntityConstraint] struct {
	Repository   Repository[Entity]
	CreateFilter CreateFilterFunction // ajout d'un champ pour stocker une fonction de génération de filtre
}

func (controller Controller[Entity]) ReadAll(res http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query())
	var query = r.URL.Query()
	filter := bson.D{}

	for k, values := range query {
		if len(values) > 0 {
			filter = append(filter, bson.E{Key: k, Value: values[0]})
		}
	}
	fmt.Println(filter)
	results, _ := controller.Repository.FindAll(filter)
	send.Json(res, results)
}

func (controller Controller[Entity]) Create(res http.ResponseWriter, r *http.Request) {
	var message Entity
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	err = controller.Repository.InsertOne(message)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	send.Json(res, err)
}

func (controller Controller[Entity]) Read(res http.ResponseWriter, r *http.Request) {
	result, err := controller.Repository.FindById(r.PathValue("id"))
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	send.Json(res, result)
}

func (controller Controller[Entity]) Delete(res http.ResponseWriter, r *http.Request) {
	result, err := controller.Repository.Delete(r.PathValue("id"))
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	send.Json(res, result)
}
