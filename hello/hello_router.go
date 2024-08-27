package hello

import (
	"net/http"
)

// retourne un retour configuré avec les routes associées au préfixe /api/hello (mais sans le préfixe)
func Router() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("GET /world", HelloWorld)
	router.HandleFunc("GET /square/{value}", Square)
	router.HandleFunc("GET /", Controller.ReadAll)
	router.HandleFunc("GET /{id}", Controller.Read)
	router.HandleFunc("POST /", Controller.Create)
	router.HandleFunc("DELETE /{id}", Controller.Delete)
	return router
}
