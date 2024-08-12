package transactionsvc

import (
	"context"

	transactionmodel "github.com/amnestia/xyz-multifinance/internal/domain/model/transaction"
	"github.com/amnestia/xyz-multifinance/internal/lib/account"
	libhmac "github.com/amnestia/xyz-multifinance/internal/lib/crypto/hmac"
)

var (
	getPartnerData   = account.GetPartnerData
	generateHMACHash = libhmac.GetHash
)

var (
	checkLimit = func(svc *Service, ctx context.Context, payload transactionmodel.LookupPayload, req *transactionmodel.TransactionRequest) (err error) {
		return svc.checkLimit(ctx, payload, req)
	}
)
