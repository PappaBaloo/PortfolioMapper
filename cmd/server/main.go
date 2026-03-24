package main

import (
	"log"
	"net/http"

	"github.com/PappaBaloo/PortfolioMapper/internal/api"
	"github.com/PappaBaloo/PortfolioMapper/internal/store"
)

type MockStore struct{}

func main() {
	log.Println("Server is running...")
	s := &store.MockStore{}
	r := api.NewRouter(s)
	http.ListenAndServe(":8080", r)
}
