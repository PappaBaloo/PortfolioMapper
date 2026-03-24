package main

import (
	"log"
	"net/http"

	"github.com/PappaBaloo/PortfolioMapper/internal/api"
)

func main() {
	log.Println("Server is running...")
	r := api.NewRouter()
	r.Get("/portfolio", func(w http.ResponseWriter, r *http.Request) {

	})
	http.ListenAndServe(":8080", r)
}
