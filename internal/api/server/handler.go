package server

import (
	"github.com/amnestia/xyz-multifinance/internal/config"
	"github.com/amnestia/xyz-multifinance/internal/database"
	"github.com/amnestia/xyz-multifinance/internal/domain/controller"
	"github.com/amnestia/xyz-multifinance/internal/domain/controller/auth"
	"github.com/amnestia/xyz-multifinance/internal/domain/controller/ping"
	"github.com/amnestia/xyz-multifinance/internal/domain/controller/transaction"
	authrepo "github.com/amnestia/xyz-multifinance/internal/domain/repository/auth"
	transactionrepo "github.com/amnestia/xyz-multifinance/internal/domain/repository/transaction"
	authsvc "github.com/amnestia/xyz-multifinance/internal/domain/service/auth"
	pingsvc "github.com/amnestia/xyz-multifinance/internal/domain/service/ping"
	transactionsvc "github.com/amnestia/xyz-multifinance/internal/domain/service/transaction"
	"github.com/amnestia/xyz-multifinance/internal/lib/paseto"
)

type dependency struct {
	db       *database.Base
	authRepo *authrepo.Repository

	past *paseto.PASTHandle
}

func getController(cfg config.Config, dep *dependency) *controller.Controller {
	pingSvc := &pingsvc.Service{}
	pingCtrl := ping.Controller{
		PingSvc: pingSvc,
	}

	// initialize repo
	authRepo := &authrepo.Repository{}
	authRepo.DB = dep.db
	dep.authRepo = authRepo

	transactionRepo := &transactionrepo.Repository{}
	transactionRepo.DB = dep.db

	// initialize service
	authSvc := &authsvc.Service{
		Config: cfg,
		Repo:   authRepo,
		Paseto: dep.past,
	}

	transactionSvc := &transactionsvc.Service{
		Config:   cfg,
		Repo:     transactionRepo,
		AuthRepo: authRepo,
	}

	// initialize controller
	authCtrl := auth.Controller{
		AuthSvc: authSvc,
	}
	transactionCtrl := transaction.Controller{
		TransactionSvc: transactionSvc,
	}

	ctrl := controller.Controller{
		PingHandler:        pingCtrl,
		AuthHandler:        authCtrl,
		TransactionHandler: transactionCtrl,
	}

	return &ctrl
}
