package game

import (
	"net/http"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /", Controller.ReadAll)
	router.HandleFunc("GET /{id}", Controller.Read)
	router.HandleFunc("DELETE /{id}", Controller.Delete)
	router.HandleFunc("POST /new", NewGame)
	return router
}
