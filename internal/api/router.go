package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/PappaBaloo/PortfolioMapper/internal/store"
	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/neighbour", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Neighbour!")
	})
	router.Get("/portfolio", func(w http.ResponseWriter, r *http.Request) {
		portfolio := store.MockPortfolio()
		response := fmt.Sprintf(`{"id":"%s","name":"%s","description":"%s"}`, portfolio.ID, portfolio.Name, portfolio.Description)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(response))
	})
	return router
}
