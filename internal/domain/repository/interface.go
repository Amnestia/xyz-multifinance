package repository

import (
	"context"

	authmodel "github.com/amnestia/xyz-multifinance/internal/domain/model/auth"
	transactionmodel "github.com/amnestia/xyz-multifinance/internal/domain/model/transaction"
	"github.com/jmoiron/sqlx"
)

// Repositorier interface
type Repositorier interface {
	NewTransaction(ctx context.Context) (*sqlx.Tx, error)
	Commit(tx *sqlx.Tx) (err error)
	RollbackOnError(tx *sqlx.Tx, err error) error
}

// AuthRepository interface
type AuthRepository interface {
	Repositorier
	Auth(ctx context.Context, email string) (*authmodel.Account, error)
	RegisterNewAccount(ctx context.Context, tx *sqlx.Tx, acc *authmodel.Account) (id int64, err error)
	GetPartner(ctx context.Context, clientID string) (*authmodel.Partner, error)
}

// TransactionRepository interface
type TransactionRepository interface {
	Repositorier

	CreateNewTransaction(ctx context.Context, tx *sqlx.Tx, req transactionmodel.Transaction) (ret int64, err error)
	CreateNewPayment(ctx context.Context, tx *sqlx.Tx, req transactionmodel.Payment) (ret int64, err error)
	CreateNewPaymentInstallment(ctx context.Context, tx *sqlx.Tx, req transactionmodel.PaymentInstallment) (ret int64, err error)

	GetLimit(ctx context.Context, req transactionmodel.LookupPayload) (ret transactionmodel.Limit, err error)
	GetOngoingPayment(ctx context.Context, req transactionmodel.LookupPayload) (ret []transactionmodel.Payment, err error)
	GetOngoingPaymentInstallment(ctx context.Context, req transactionmodel.LookupPayload) (ret []transactionmodel.PaymentInstallment, err error)
}
