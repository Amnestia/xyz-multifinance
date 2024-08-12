package transactionsvc

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"testing"

	"github.com/amnestia/xyz-multifinance/internal/domain/constant"
	authmodel "github.com/amnestia/xyz-multifinance/internal/domain/model/auth"
	transactionmodel "github.com/amnestia/xyz-multifinance/internal/domain/model/transaction"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewTransaction(t *testing.T) {
	var (
		ctx             = context.Background()
		mockErr         = errors.New("mock error")
		tx              = &sqlx.Tx{}
		mockID          = int64(1)
		mockReq         = &transactionmodel.TransactionRequest{}
		mockAcc         = &authmodel.Account{}
		mockPartnerData = &authmodel.Partner{}
		mockTransaction = &transactionmodel.Transaction{}
		mockPayment     = &transactionmodel.Payment{
			TransactionID: 1,
			Status:        constant.PaymentUnpaidStatus,
		}
		mockPaymentInstallment = &transactionmodel.PaymentInstallment{
			PaymentID: 1,
			Status:    constant.PaymentUnpaidStatus,
		}

		tmpGetPartnerData   = getPartnerData
		tmpGenerateHMACHash = generateHMACHash
		tmpCheckLimit       = checkLimit
	)
	defer func() {
		getPartnerData = tmpGetPartnerData
		generateHMACHash = tmpGenerateHMACHash
		checkLimit = tmpCheckLimit
	}()
	type args struct {
		ctx context.Context
		req *transactionmodel.TransactionRequest
	}
	type expect struct {
		httpCode int
		err      bool
	}
	tests := []struct {
		name   string
		args   args
		expect expect
		patch  func(args)
	}{
		{
			name: "CreateNewTransaction_ErrorOnGetPartnerData",
			args: args{
				ctx: ctx,
				req: mockReq,
			},

			expect: expect{
				httpCode: http.StatusInternalServerError,
				err:      true,
			},
			patch: func(a args) {
				getPartnerData = func(ctx context.Context) (*authmodel.Partner, error) {
					return nil, mockErr
				}
			},
		},
		{
			name: "CreateNewTransaction_ErrorOnGenerateHMACHashForNIK",
			args: args{
				ctx: ctx,
				req: mockReq,
			},

			expect: expect{
				httpCode: http.StatusInternalServerError,
				err:      true,
			},
			patch: func(a args) {
				getPartnerData = func(ctx context.Context) (*authmodel.Partner, error) {
					return mockPartnerData, nil
				}
				generateHMACHash = func(target, key, salt string) (ret string, err error) {
					return "", mockErr
				}
			},
		},
		{
			name: "CreateNewTransaction_ErrorOnAccountNotFound",
			args: args{
				ctx: ctx,
				req: mockReq,
			},

			expect: expect{
				httpCode: http.StatusNotFound,
				err:      true,
			},
			patch: func(a args) {
				getPartnerData = func(ctx context.Context) (*authmodel.Partner, error) {
					return mockPartnerData, nil
				}
				generateHMACHash = func(target, key, salt string) (ret string, err error) {
					return "", nil
				}
				mockAuthRepo.On("Auth", ctx, "").Return(mockAcc, sql.ErrNoRows).Once()
			},
		},
		{
			name: "CreateNewTransaction_ErrorOnAuthGetAccount",
			args: args{
				ctx: ctx,
				req: mockReq,
			},

			expect: expect{
				httpCode: http.StatusInternalServerError,
				err:      true,
			},
			patch: func(a args) {
				getPartnerData = func(ctx context.Context) (*authmodel.Partner, error) {
					return mockPartnerData, nil
				}
				generateHMACHash = func(target, key, salt string) (ret string, err error) {
					return "", nil
				}
				mockAuthRepo.On("Auth", ctx, "").Return(mockAcc, mockErr).Once()
			},
		},
		{
			name: "CreateNewTransaction_ErrorOnLimitReached",
			args: args{
				ctx: ctx,
				req: mockReq,
			},

			expect: expect{
				httpCode: http.StatusBadRequest,
				err:      true,
			},
			patch: func(a args) {
				getPartnerData = func(ctx context.Context) (*authmodel.Partner, error) {
					return mockPartnerData, nil
				}
				generateHMACHash = func(target, key, salt string) (ret string, err error) {
					return "", nil
				}
				mockAuthRepo.On("Auth", ctx, "").Return(mockAcc, nil).Once()
				checkLimit = func(svc *Service, ctx context.Context, payload transactionmodel.LookupPayload, req *transactionmodel.TransactionRequest) (err error) {
					return constant.OverlimitError{}
				}
			},
		},
		{
			name: "CreateNewTransaction_ErrorOnCheckLimit",
			args: args{
				ctx: ctx,
				req: mockReq,
			},

			expect: expect{
				httpCode: http.StatusInternalServerError,
				err:      true,
			},
			patch: func(a args) {
				getPartnerData = func(ctx context.Context) (*authmodel.Partner, error) {
					return mockPartnerData, nil
				}
				generateHMACHash = func(target, key, salt string) (ret string, err error) {
					return "", nil
				}
				mockAuthRepo.On("Auth", ctx, "").Return(mockAcc, nil).Once()
				checkLimit = func(svc *Service, ctx context.Context, payload transactionmodel.LookupPayload, req *transactionmodel.TransactionRequest) (err error) {
					return mockErr
				}
			},
		},
		{
			name: "CreateNewTransaction_ErrorOnStartNewDBTransaction",
			args: args{
				ctx: ctx,
				req: mockReq,
			},

			expect: expect{
				httpCode: http.StatusInternalServerError,
				err:      true,
			},
			patch: func(a args) {
				getPartnerData = func(ctx context.Context) (*authmodel.Partner, error) {
					return mockPartnerData, nil
				}
				generateHMACHash = func(target, key, salt string) (ret string, err error) {
					return "", nil
				}
				mockAuthRepo.On("Auth", ctx, "").Return(mockAcc, nil).Once()
				checkLimit = func(svc *Service, ctx context.Context, payload transactionmodel.LookupPayload, req *transactionmodel.TransactionRequest) (err error) {
					return nil
				}
				mockRepo.On("NewTransaction", ctx).Return(tx, mockErr).Once()
			},
		},
		{
			name: "CreateNewTransaction_ErrorOnCreateNewTransaction",
			args: args{
				ctx: ctx,
				req: mockReq,
			},

			expect: expect{
				httpCode: http.StatusInternalServerError,
				err:      true,
			},
			patch: func(a args) {
				getPartnerData = func(ctx context.Context) (*authmodel.Partner, error) {
					return mockPartnerData, nil
				}
				generateHMACHash = func(target, key, salt string) (ret string, err error) {
					return "", nil
				}
				mockAuthRepo.On("Auth", ctx, "").Return(mockAcc, nil).Once()
				checkLimit = func(svc *Service, ctx context.Context, payload transactionmodel.LookupPayload, req *transactionmodel.TransactionRequest) (err error) {
					return nil
				}
				mockRepo.On("NewTransaction", ctx).Return(tx, nil).Once()
				mockRepo.On("CreateNewTransaction", ctx, tx, mockTransaction).Return(mockID, mockErr).Once()
				mockRepo.On("RollbackOnError", tx, mockErr).Return(mockErr).Once()
			},
		},
		{
			name: "CreateNewTransaction_ErrorOnCreateNewPayment",
			args: args{
				ctx: ctx,
				req: mockReq,
			},

			expect: expect{
				httpCode: http.StatusInternalServerError,
				err:      true,
			},
			patch: func(a args) {
				getPartnerData = func(ctx context.Context) (*authmodel.Partner, error) {
					return mockPartnerData, nil
				}
				generateHMACHash = func(target, key, salt string) (ret string, err error) {
					return "", nil
				}
				mockAuthRepo.On("Auth", ctx, "").Return(mockAcc, nil).Once()
				checkLimit = func(svc *Service, ctx context.Context, payload transactionmodel.LookupPayload, req *transactionmodel.TransactionRequest) (err error) {
					return nil
				}
				mockRepo.On("NewTransaction", ctx).Return(tx, nil).Once()
				mockRepo.On("CreateNewTransaction", ctx, tx, mockTransaction).Return(mockID, nil).Once()
				mockRepo.On("RollbackOnError", tx, nil).Return(nil).Once()
				mockRepo.On("CreateNewPayment", ctx, tx, mockPayment).Return(mockID, mockErr).Once()
				mockRepo.On("RollbackOnError", tx, mockErr).Return(mockErr).Once()
			},
		},
		{
			name: "CreateNewTransaction_ErrorOnCreateNewPaymentInstallment",
			args: args{
				ctx: ctx,
				req: mockReq,
			},

			expect: expect{
				httpCode: http.StatusInternalServerError,
				err:      true,
			},
			patch: func(a args) {
				getPartnerData = func(ctx context.Context) (*authmodel.Partner, error) {
					return mockPartnerData, nil
				}
				generateHMACHash = func(target, key, salt string) (ret string, err error) {
					return "", nil
				}
				mockAuthRepo.On("Auth", ctx, "").Return(mockAcc, nil).Once()
				checkLimit = func(svc *Service, ctx context.Context, payload transactionmodel.LookupPayload, req *transactionmodel.TransactionRequest) (err error) {
					return nil
				}
				mockRepo.On("NewTransaction", ctx).Return(tx, nil).Once()
				mockRepo.On("CreateNewTransaction", ctx, tx, mockTransaction).Return(mockID, nil).Once()
				mockRepo.On("RollbackOnError", tx, nil).Return(nil).Once()
				mockRepo.On("CreateNewPayment", ctx, tx, mockPayment).Return(mockID, nil).Once()
				mockRepo.On("RollbackOnError", tx, nil).Return(nil).Once()
				mockRepo.On("CreateNewPaymentInstallment", ctx, tx, mockPaymentInstallment).Return(mockID, mockErr).Once()
				mockRepo.On("RollbackOnError", tx, mockErr).Return(mockErr).Once()
			},
		},
		{
			name: "CreateNewTransaction_ErrorOnCommit",
			args: args{
				ctx: ctx,
				req: mockReq,
			},

			expect: expect{
				httpCode: http.StatusInternalServerError,
				err:      true,
			},
			patch: func(a args) {
				getPartnerData = func(ctx context.Context) (*authmodel.Partner, error) {
					return mockPartnerData, nil
				}
				generateHMACHash = func(target, key, salt string) (ret string, err error) {
					return "", nil
				}
				mockAuthRepo.On("Auth", ctx, "").Return(mockAcc, nil).Once()
				checkLimit = func(svc *Service, ctx context.Context, payload transactionmodel.LookupPayload, req *transactionmodel.TransactionRequest) (err error) {
					return nil
				}
				mockRepo.On("NewTransaction", ctx).Return(tx, nil).Once()
				mockRepo.On("CreateNewTransaction", ctx, tx, mockTransaction).Return(mockID, nil).Once()
				mockRepo.On("RollbackOnError", tx, nil).Return(nil).Once()
				mockRepo.On("CreateNewPayment", ctx, tx, mockPayment).Return(mockID, nil).Once()
				mockRepo.On("RollbackOnError", tx, nil).Return(nil).Once()
				mockRepo.On("CreateNewPaymentInstallment", ctx, tx, mockPaymentInstallment).Return(mockID, nil).Once()
				mockRepo.On("RollbackOnError", tx, nil).Return(nil).Once()
				mockRepo.On("Commit", tx).Return(mockErr).Once()
			},
		},
		{
			name: "CreateNewTransaction_NoError",
			args: args{
				ctx: ctx,
				req: mockReq,
			},

			expect: expect{
				httpCode: http.StatusCreated,
				err:      false,
			},
			patch: func(a args) {
				getPartnerData = func(ctx context.Context) (*authmodel.Partner, error) {
					return mockPartnerData, nil
				}
				generateHMACHash = func(target, key, salt string) (ret string, err error) {
					return "", nil
				}
				mockAuthRepo.On("Auth", ctx, "").Return(mockAcc, nil).Once()
				checkLimit = func(svc *Service, ctx context.Context, payload transactionmodel.LookupPayload, req *transactionmodel.TransactionRequest) (err error) {
					return nil
				}
				mockRepo.On("NewTransaction", ctx).Return(tx, nil).Once()
				mockRepo.On("CreateNewTransaction", ctx, tx, mockTransaction).Return(mockID, nil).Once()
				mockRepo.On("RollbackOnError", tx, nil).Return(nil).Once()
				mockRepo.On("CreateNewPayment", ctx, tx, mockPayment).Return(mockID, nil).Once()
				mockRepo.On("RollbackOnError", tx, nil).Return(nil).Once()
				mockRepo.On("CreateNewPaymentInstallment", ctx, tx, mockPaymentInstallment).Return(mockID, nil).Once()
				mockRepo.On("RollbackOnError", tx, nil).Return(nil).Once()
				mockRepo.On("Commit", tx).Return(nil).Once()
			},
		},
	}
	for _, test := range tests {
		svc := getMock()
		test.patch(test.args)
		t.Run(test.name, func(t *testing.T) {
			resp := svc.CreateNewTransaction(test.args.ctx, test.args.req)
			assert.Equal(t, test.expect.httpCode, resp.HTTPCode)
			if test.expect.err {
				assert.Error(t, resp.Error)
			} else {
				assert.NoError(t, resp.Error)
			}
		})
	}
}
