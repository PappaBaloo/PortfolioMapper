package store

import (
	"github.com/PappaBaloo/PortfolioMapper/internal/models"
)

type Store interface {
	GetPortfolio() *models.Portfolio
}
