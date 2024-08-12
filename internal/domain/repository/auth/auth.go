package authrepo

import (
	"context"
	"database/sql"

	authmodel "github.com/amnestia/xyz-multifinance/internal/domain/model/auth"
	"github.com/amnestia/xyz-multifinance/internal/domain/repository"
	"github.com/amnestia/xyz-multifinance/pkg/logger"
	"github.com/jmoiron/sqlx"
)

// Repository struct
type Repository struct {
	repository.Repository
}

// Auth get credentials from database
func (repo *Repository) Auth(ctx context.Context, nik string) (*authmodel.Account, error) {
	acc := authmodel.Account{}
	err := repo.DB.Slave.GetContext(ctx, &acc, auth, nik)
	if err != nil {
		if err == sql.ErrNoRows {
			return &acc, err
		}
		return &acc, logger.ErrorWrap(err, "Auth", "GetContext")
	}
	return &acc, nil
}

// RegisterNewAccount register new account to database
func (repo *Repository) RegisterNewAccount(ctx context.Context, tx *sqlx.Tx, acc *authmodel.Account) (id int64, err error) {
	q, args, err := sqlx.Named(insertNewAccount, &acc)
	if err != nil {
		return -1, logger.ErrorWrap(err, "RegisterNewAccount", "SQLXNamed")
	}
	q = repo.DB.Slave.Rebind(q)
	res, err := tx.ExecContext(ctx, q, args...)
	if err != nil {
		return -1, logger.ErrorWrap(err, "RegisterNewAccount", "ExecContext")
	}
	id, err = res.LastInsertId()
	if err != nil {
		return -1, logger.ErrorWrap(err, "RegisterNewAccount", "LastInsertId")
	}
	return
}

// GetPartner get partner credentials by client id
func (repo *Repository) GetPartner(ctx context.Context, clientID string) (*authmodel.Partner, error) {
	acc := authmodel.Partner{}
	err := repo.DB.Slave.GetContext(ctx, &acc, getPartner, clientID)
	if err != nil {
		if err == sql.ErrNoRows {
			return &acc, err
		}
		return &acc, logger.ErrorWrap(err, "GetPartner.GetContext")
	}
	return &acc, nil
}
