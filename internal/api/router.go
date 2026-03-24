package api

import (
	"io"
	"net/http"

	"encoding/json"

	"github.com/PappaBaloo/PortfolioMapper/internal/store"
	"github.com/go-chi/chi/v5"
)

func NewRouter(store store.Store) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/neighbour", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Neighbour!")
	})
	router.Get("/portfolio", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		portfolio := store.GetPortfolio()
		json.NewEncoder(w).Encode(portfolio)
	})
	return router
}
