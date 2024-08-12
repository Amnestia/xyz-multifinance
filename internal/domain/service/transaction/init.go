package transactionsvc

import (
	"github.com/amnestia/xyz-multifinance/internal/config"
	"github.com/amnestia/xyz-multifinance/internal/domain/repository"
)

// Service service functionality of the domain
type Service struct {
	Config   config.Config
	Repo     repository.TransactionRepository
	AuthRepo repository.AuthRepository
}
