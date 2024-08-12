package transactionsvc

import (
	"context"
	"database/sql"

	"github.com/amnestia/xyz-multifinance/internal/domain/constant"
	transactionmodel "github.com/amnestia/xyz-multifinance/internal/domain/model/transaction"
	"github.com/amnestia/xyz-multifinance/pkg/logger"
)

func (svc *Service) checkLimit(ctx context.Context, payload transactionmodel.LookupPayload, req *transactionmodel.TransactionRequest) (err error) {
	limit, err := svc.Repo.GetLimit(ctx, payload)
	if err != nil {
		err = logger.ErrorWrap(err, "CreateNewTransaction.CheckLimit")
		return
	}
	payments, err := svc.Repo.GetOngoingPayment(ctx, payload)
	if err != nil && err != sql.ErrNoRows {
		err = logger.ErrorWrap(err, "CreateNewTransaction.GetOngoingPayment")
		return
	}
	currentLimit := limit.Amount
	if len(payments) > 0 {
		for _, val := range payments {
			payload.PaymentID = append(payload.PaymentID, val.ID)
		}
		unpaidInstallments, err := svc.Repo.GetOngoingPaymentInstallment(ctx, payload)
		if err != nil && err != sql.ErrNoRows {
			err = logger.ErrorWrap(err, "CreateNewTransaction.GetOngoingPaymentInstallment")
			return err
		}
		for _, val := range unpaidInstallments {
			currentLimit -= val.Amount
		}
	}
	if currentLimit < req.TotalInstallment {
		return constant.OverlimitError{}
	}
	return nil
}
