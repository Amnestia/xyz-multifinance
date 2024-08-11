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
func (repo *Repository) Auth(ctx context.Context, email string) (*authmodel.Account, error) {
	acc := &authmodel.Account{}
	err := repo.DB.Slave.QueryRowxContext(ctx, auth, email).Scan(&acc.Email, &acc.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return acc, err
		}
		return acc, logger.ErrorWrap(err, "repo", "Failed on getting account")
	}
	return acc, nil
}

// RegisterNewAccount register new account to database
func (repo *Repository) RegisterNewAccount(ctx context.Context, tx *sqlx.Tx, acc *authmodel.Account) (id int64, err error) {
	err = tx.GetContext(ctx, &id, insertNewAccount, &acc)
	if err != nil {
		return -1, logger.ErrorWrap(err, "repo", "Failed on creating new account")
	}
	return
}
