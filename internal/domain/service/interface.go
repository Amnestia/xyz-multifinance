package service

import (
	"context"

	authmodel "github.com/amnestia/xyz-multifinance/internal/domain/model/auth"
	"github.com/amnestia/xyz-multifinance/internal/domain/model/common"
	transactionmodel "github.com/amnestia/xyz-multifinance/internal/domain/model/transaction"
)

// PingServicer interface
type PingServicer interface {
	Ping() string
}

// AuthServicer interface
type AuthServicer interface {
	Register(context.Context, *authmodel.RegisterRequest) *common.DefaultResponse
	Auth(context.Context, *authmodel.ConsumerAuthRequest) *authmodel.LoginResponse

	CreateNewPartner(ctx context.Context, req *authmodel.Partner) (resp *authmodel.NewPartnerResponse)
}

// TransactionServicer interface
type TransactionServicer interface {
	CreateNewTransaction(context.Context, *transactionmodel.TransactionRequest) *common.DefaultResponse
}
