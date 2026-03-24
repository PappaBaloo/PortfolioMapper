package store

import (
	"github.com/PappaBaloo/PortfolioMapper/internal/models"
)

func mockPortfolio() *models.Portfolio {
	return &models.Portfolio{
		ID:          "1",
		Name:        "My Portfolio",
		Description: "A simple portfolio",
	}
}
