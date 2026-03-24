package store

import (
	"github.com/PappaBaloo/PortfolioMapper/internal/models"
)

type MockStore struct {
}

func (m *MockStore) GetPortfolio() *models.Portfolio {
	return &models.Portfolio{
		ID:          "1",
		Name:        "Portfolio 1",
		Description: "A simple portfolio",
	}
}
