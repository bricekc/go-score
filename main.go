package main

import (
	"but3/go-score/game"
	"but3/go-score/hello"
	"but3/go-score/player"
	"but3/go-score/services/mongodb"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	if mongodb.Collection("hello").Name() != "hello" {
		panic("Problème d'accès à la collection hello")
	}
	if mongodb.Collection("player").Name() != "player" {
		panic("Problème d'accès à la collection player")
	}
	if mongodb.Collection("game").Name() != "game" {
		panic("Problème d'accès à la collection game")
	}
	// création du routeur
	router := http.NewServeMux()
	router.Handle("/api/hello/", http.StripPrefix("/api/hello", hello.Router()))

	router.Handle("/api/player/", http.StripPrefix("/api/player", player.Router()))

	router.Handle("/api/game/", http.StripPrefix("/api/game", game.Router()))

	// configuration du serveur
	server := http.Server{
		Addr:    ":8888",
		Handler: cors.AllowAll().Handler(router),
	}

	log.Println("Serveur local démarré : http://localhost:8888")

	// démarrage du serveur
	err := server.ListenAndServe()

	log.Fatalln(err)
}
