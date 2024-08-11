package server

import (
	"github.com/amnestia/xyz-multifinance/internal/config"
	"github.com/amnestia/xyz-multifinance/internal/database"
	"github.com/amnestia/xyz-multifinance/internal/domain/controller"
	"github.com/amnestia/xyz-multifinance/internal/domain/controller/auth"
	"github.com/amnestia/xyz-multifinance/internal/domain/controller/ping"
	authrepo "github.com/amnestia/xyz-multifinance/internal/domain/repository/auth"
	authsvc "github.com/amnestia/xyz-multifinance/internal/domain/service/auth"
	pingsvc "github.com/amnestia/xyz-multifinance/internal/domain/service/ping"
	"github.com/amnestia/xyz-multifinance/internal/lib/paseto"
)

type dependency struct {
	db *database.Base

	past *paseto.PASTHandle
}

func getController(cfg config.Config, dep dependency) *controller.Controller {
	pingSvc := &pingsvc.Service{}
	pingCtrl := ping.Controller{
		PingSvc: pingSvc,
	}

	// initialize repo
	authRepo := &authrepo.Repository{}
	authRepo.DB = dep.db

	// initialize service
	authSvc := &authsvc.Service{
		Config: cfg,
		Repo:   authRepo,
		Paseto: dep.past,
	}

	// initialize controller
	authCtrl := auth.Controller{
		AuthSvc: authSvc,
	}

	ctrl := controller.Controller{
		PingHandler: pingCtrl,
		AuthHandler: authCtrl,
	}

	return &ctrl
}
