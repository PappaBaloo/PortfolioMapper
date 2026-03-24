package api

import (
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/nigger", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Nigger!")
	})
	return router
}
