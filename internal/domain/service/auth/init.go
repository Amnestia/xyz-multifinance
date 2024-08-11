package authsvc

import (
	"github.com/amnestia/xyz-multifinance/internal/config"
	"github.com/amnestia/xyz-multifinance/internal/domain/repository"
	"github.com/amnestia/xyz-multifinance/internal/lib/paseto"
)

// Service service functionality of the domain
type Service struct {
	Config config.Config
	Repo   repository.AuthRepository
	Paseto paseto.Handler
}
