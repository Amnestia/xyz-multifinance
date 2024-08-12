package transactionsvc

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/amnestia/xyz-multifinance/internal/domain/constant"
	"github.com/amnestia/xyz-multifinance/internal/domain/model/common"
	transactionmodel "github.com/amnestia/xyz-multifinance/internal/domain/model/transaction"
	"github.com/amnestia/xyz-multifinance/pkg/logger"
)

func (svc *Service) CreateNewTransaction(ctx context.Context, req *transactionmodel.TransactionRequest) (resp *common.DefaultResponse) {
	resp = &common.DefaultResponse{HTTPCode: http.StatusCreated}
	partner, err := getPartnerData(ctx)
	if err != nil {
		err = logger.ErrorWrap(err, "CreateNewTransaction.GetPartnerData")
		resp.Build(http.StatusInternalServerError, err)
		return
	}
	nik, err := generateHMACHash(req.NIK, svc.Config.Crypt.HMAC.Key, svc.Config.Crypt.HMAC.Pepper)
	if err != nil {
		err = logger.ErrorWrap(err, "CreateNewTransaction.HMACGetHash")
		resp.Build(http.StatusInternalServerError, err)
		return
	}
	acc, err := svc.AuthRepo.Auth(ctx, nik)
	if err != nil {
		if err == sql.ErrNoRows {
			resp.Build(http.StatusNotFound, constant.LoginFailedError{})
			return
		}
		err = logger.ErrorWrap(err, "CreateNewTransaction.RepoAuth")
		resp.Build(http.StatusInternalServerError, err)
		return
	}
	payload := transactionmodel.LookupPayload{
		ConsumerID: acc.ID,
		Duration:   req.Duration,
		Status:     constant.PaymentUnpaidStatus,
	}
	err = checkLimit(svc, ctx, payload, req)
	if err != nil {
		if errors.Is(err, constant.OverlimitError{}) {
			resp.Build(http.StatusBadRequest, constant.OverlimitError{})
			return
		}
		resp.Build(http.StatusInternalServerError, err)
		return
	}
	tx, err := svc.Repo.NewTransaction(ctx)
	if err != nil {
		resp = resp.Build(http.StatusInternalServerError, err)
		return
	}
	transactionID, err := svc.Repo.CreateNewTransaction(ctx, tx, &transactionmodel.Transaction{
		ContractNumber:   req.ContractNumber,
		AssetName:        req.AssetName,
		ConsumerID:       acc.ID,
		PartnerID:        partner.ID,
		OTR:              req.OTR,
		AdminFee:         req.AdminFee,
		TotalInstallment: req.TotalInstallment,
		Interest:         req.Interest,
	})
	if err = svc.Repo.RollbackOnError(tx, err); err != nil {
		resp = resp.Build(http.StatusInternalServerError, err)
		return
	}
	paymentID, err := svc.Repo.CreateNewPayment(ctx, tx, &transactionmodel.Payment{
		TransactionID: transactionID,
		ConsumerID:    acc.ID,
		TotalAmount:   req.OTR + req.AdminFee,
		MonthlyAmount: req.TotalInstallment,
		Duration:      req.Duration,
		Interest:      req.Interest,
		Status:        constant.PaymentUnpaidStatus,
	})
	if err = svc.Repo.RollbackOnError(tx, err); err != nil {
		resp = resp.Build(http.StatusInternalServerError, err)
		return
	}
	_, err = svc.Repo.CreateNewPaymentInstallment(ctx, tx, &transactionmodel.PaymentInstallment{
		PaymentID: paymentID,
		Amount:    req.TotalInstallment,
		Status:    constant.PaymentUnpaidStatus,
	})
	if err = svc.Repo.RollbackOnError(tx, err); err != nil {
		resp = resp.Build(http.StatusInternalServerError, err)
		return
	}
	err = svc.Repo.Commit(tx)
	if err != nil {
		resp = resp.Build(http.StatusInternalServerError, err)
	}
	return
}
