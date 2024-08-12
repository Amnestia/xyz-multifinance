package controller

import (
	"github.com/amnestia/xyz-multifinance/internal/domain/controller/auth"
	"github.com/amnestia/xyz-multifinance/internal/domain/controller/ping"
	"github.com/amnestia/xyz-multifinance/internal/domain/controller/transaction"
)

// Controller controller containing handler for services
type Controller struct {
	PingHandler        ping.Controller
	AuthHandler        auth.Controller
	TransactionHandler transaction.Controller
}
