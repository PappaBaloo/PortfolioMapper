package api

import (
	"io"
	"net/http"

	"encoding/json"

	"github.com/PappaBaloo/PortfolioMapper/internal/store"
	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/neighbour", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Neighbour!")
	})
	router.Get("/portfolio", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		portfolio := store.MockPortfolio()
		response := json.NewEncoder(w).Encode(portfolio)
		if response != nil {
			http.Error(w, "Failed to encode portfolio", http.StatusInternalServerError)
			return
		}
	})
	return router
}
