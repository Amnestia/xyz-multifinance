package transactionrepo

import (
	"context"
	"database/sql"

	transactionmodel "github.com/amnestia/xyz-multifinance/internal/domain/model/transaction"
	"github.com/amnestia/xyz-multifinance/internal/domain/repository"
	"github.com/amnestia/xyz-multifinance/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	repository.Repository
}

// CreateNewTransaction create new transaction
func (repo *Repository) CreateNewTransaction(ctx context.Context, tx *sqlx.Tx, req *transactionmodel.Transaction) (id int64, err error) {
	q, args, err := sqlx.Named(insertNewTransaction, &req)
	if err != nil {
		return -1, logger.ErrorWrap(err, "CreateNewTransaction.SQLXNamed")
	}
	q = repo.DB.Slave.Rebind(q)
	res, err := tx.ExecContext(ctx, q, args...)
	if err != nil {
		return -1, logger.ErrorWrap(err, "CreateNewTransaction.SQLXNamed")
	}
	id, err = res.LastInsertId()
	if err != nil {
		return -1, logger.ErrorWrap(err, "CreateNewTransaction.SQLXNamed")
	}
	return
}

// CreateNewPayment create new payment
func (repo *Repository) CreateNewPayment(ctx context.Context, tx *sqlx.Tx, req *transactionmodel.Payment) (id int64, err error) {
	q, args, err := sqlx.Named(insertNewPayment, &req)
	if err != nil {
		return -1, logger.ErrorWrap(err, "CreateNewPayment.SQLXNamed")
	}
	q = repo.DB.Slave.Rebind(q)
	res, err := tx.ExecContext(ctx, q, args...)
	if err != nil {
		return -1, logger.ErrorWrap(err, "CreateNewPayment.ExecContext")
	}
	id, err = res.LastInsertId()
	if err != nil {
		return -1, logger.ErrorWrap(err, "CreateNewPayment.LastInsertId")
	}
	return
}

// CreateNewPaymentInstallment create new payment installment
func (repo *Repository) CreateNewPaymentInstallment(ctx context.Context, tx *sqlx.Tx, req *transactionmodel.PaymentInstallment) (id int64, err error) {
	q, args, err := sqlx.Named(insertNewPaymentInstallment, &req)
	if err != nil {
		return -1, logger.ErrorWrap(err, "CreateNewPaymentInstallment.SQLXNamed")
	}
	q = repo.DB.Slave.Rebind(q)
	res, err := tx.ExecContext(ctx, q, args...)
	if err != nil {
		return -1, logger.ErrorWrap(err, "CreateNewPaymentInstallment.ExecContext")
	}
	id, err = res.LastInsertId()
	if err != nil {
		return -1, logger.ErrorWrap(err, "CreateNewPaymentInstallment.LastInsertId")
	}
	return
}

func (repo *Repository) GetLimit(ctx context.Context, req transactionmodel.LookupPayload) (ret transactionmodel.Limit, err error) {
	err = repo.DB.Slave.GetContext(ctx, &ret, getLimit, req.ConsumerID, req.Duration)
	if err != nil {
		if err == sql.ErrNoRows {
			return ret, err
		}
		err = logger.ErrorWrap(err, "getOngoingPayment")
		return
	}
	return
}

func (repo *Repository) GetOngoingPayment(ctx context.Context, req transactionmodel.LookupPayload) (ret []transactionmodel.Payment, err error) {
	err = repo.DB.Slave.SelectContext(ctx, &ret, getOngoingPayment, req.ConsumerID, req.Duration, req.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return ret, err
		}
		err = logger.ErrorWrap(err, "getOngoingPayment")
		return
	}
	return
}

func (repo *Repository) GetOngoingPaymentInstallment(ctx context.Context, req transactionmodel.LookupPayload) (ret []transactionmodel.PaymentInstallment, err error) {
	q, args, err := sqlx.Named(getOngoingPaymentInstallment, map[string]interface{}{
		"consumer_id": req.ConsumerID,
		"status":      req.Status,
		"payment_id":  req.PaymentID,
	})
	if err != nil {
		err = logger.ErrorWrap(err, "getOngoingPayment.SQLXNamed")
		return
	}
	q, args, err = sqlx.In(q, args...)
	if err != nil {
		err = logger.ErrorWrap(err, "getOngoingPayment.SQLXIn")
		return
	}
	q = repo.DB.Slave.Rebind(q)
	err = repo.DB.Slave.SelectContext(ctx, &ret, q, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return ret, err
		}
		err = logger.ErrorWrap(err, "getOngoingPayment.SelectContext")
		return
	}
	return
}
