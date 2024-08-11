package authsvc

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/amnestia/xyz-multifinance/internal/domain/constant"
	authmodel "github.com/amnestia/xyz-multifinance/internal/domain/model/auth"
	"github.com/amnestia/xyz-multifinance/internal/domain/model/common"
	"github.com/amnestia/xyz-multifinance/internal/lib/paseto"
)

// Register register new account
func (svc *Service) Register(ctx context.Context, req *authmodel.RegisterRequest) (resp *common.DefaultResponse) {
	var err error
	resp = &common.DefaultResponse{HTTPCode: http.StatusCreated}
	acc, err := buildConsumerRegistrationData(svc, ctx, req)
	if err != nil {
		resp = resp.Build(http.StatusInternalServerError, err)
		return
	}
	tx, err := svc.Repo.NewTransaction(ctx)
	if err != nil {
		resp = resp.Build(http.StatusInternalServerError, err)
		return
	}
	_, err = svc.Repo.RegisterNewAccount(ctx, tx, acc)
	if err = svc.Repo.RollbackOnError(tx, err); err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			resp = resp.Build(http.StatusBadRequest, fmt.Errorf("Email have been registered"))
			return
		}
		resp = resp.Build(http.StatusInternalServerError, err)
		return
	}
	err = svc.Repo.Commit(tx)
	if err != nil {
		resp = resp.Build(http.StatusInternalServerError, err)
	}
	return
}

// Auth authenticate account login
func (svc *Service) Auth(ctx context.Context, req *authmodel.ConsumerAuthRequest) (resp *authmodel.LoginResponse) {
	resp = &authmodel.LoginResponse{}
	resp.HTTPCode = http.StatusOK
	acc, err := getAccount(svc, ctx, req)
	if err != nil {
		resp.Build(http.StatusInternalServerError, err)
		if err == sql.ErrNoRows {
			resp.Build(http.StatusNotFound, constant.LoginFailedError{})
		}
		return
	}
	valid, err := verifyArgonHash(req.Password, acc.Password)
	if !valid || err != nil {
		resp.Build(http.StatusNotFound, constant.LoginFailedError{})
		return
	}

	payload := paseto.Payload{
		ID: acc.ID,
	}
	payload.TokenType = paseto.AccessToken
	resp.AccessToken, err = svc.Paseto.Generate(payload)
	if err != nil {
		resp.Build(http.StatusInternalServerError, err)
		return
	}
	payload.TokenType = paseto.RefreshToken
	resp.RefreshToken, err = svc.Paseto.Generate(payload)
	if err != nil {
		resp.Build(http.StatusInternalServerError, err)
		return
	}

	return
}
