package player

import (
	"net/http"
)

// retourne un retour configuré avec les routes associées au préfixe /api/hello (mais sans le préfixe)
func Router() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /", Controller.ReadAll)
	router.HandleFunc("GET /{id}", Controller.Read)
	router.HandleFunc("POST /", Controller.Create)
	router.HandleFunc("DELETE /{id}", Controller.Delete)
	return router
}
