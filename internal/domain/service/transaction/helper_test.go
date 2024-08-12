package transactionsvc

import (
	"context"
	"errors"
	"testing"

	transactionmodel "github.com/amnestia/xyz-multifinance/internal/domain/model/transaction"
	"github.com/stretchr/testify/assert"
)

func TestCheckLimit(t *testing.T) {
	var (
		ctx         = context.Background()
		mockErr     = errors.New("mock error")
		mockPayload = transactionmodel.LookupPayload{}
		mockReq     = &transactionmodel.TransactionRequest{
			TotalInstallment: 1000,
		}
		mockEnoughLimit        = transactionmodel.Limit{Amount: 1000}
		mockLimit              = transactionmodel.Limit{}
		mockPayment            = []transactionmodel.Payment{{}, {}}
		mockPaymentInstallment = []transactionmodel.PaymentInstallment{{}, {}}
		mockInstallmentPayload = transactionmodel.LookupPayload{
			PaymentID: []int64{0, 0},
		}
	)
	type args struct {
		ctx     context.Context
		payload transactionmodel.LookupPayload
		req     *transactionmodel.TransactionRequest
	}
	type expect struct {
		err bool
	}
	tests := []struct {
		name   string
		args   args
		expect expect
		patch  func(args)
	}{
		{
			name: "checkLimit_ErrorOnGetLimit",
			args: args{
				ctx:     ctx,
				payload: mockPayload,
				req:     mockReq,
			},
			expect: expect{
				err: true,
			},
			patch: func(a args) {
				mockRepo.On("GetLimit", ctx, mockPayload).Return(mockLimit, mockErr).Once()
			},
		},
		{
			name: "checkLimit_ErrorOnGetOngoingPayment",
			args: args{
				ctx:     ctx,
				payload: mockPayload,
				req:     mockReq,
			},
			expect: expect{
				err: true,
			},
			patch: func(a args) {
				mockRepo.On("GetLimit", ctx, mockPayload).Return(mockLimit, nil).Once()
				mockRepo.On("GetOngoingPayment", ctx, mockPayload).Return(nil, mockErr).Once()
			},
		},
		{
			name: "checkLimit_ErrorOnGetOngoingPaymentInstallment",
			args: args{
				ctx:     ctx,
				payload: mockPayload,
				req:     mockReq,
			},
			expect: expect{
				err: true,
			},
			patch: func(a args) {
				mockRepo.On("GetLimit", ctx, mockPayload).Return(mockLimit, nil).Once()
				mockRepo.On("GetOngoingPayment", ctx, mockPayload).Return(mockPayment, nil).Once()
				mockRepo.On("GetOngoingPaymentInstallment", ctx, mockInstallmentPayload).Return(nil, mockErr).Once()
			},
		},
		{
			name: "checkLimit_ErrorLimitIsNotEnough",
			args: args{
				ctx:     ctx,
				payload: mockPayload,
				req:     mockReq,
			},
			expect: expect{
				err: true,
			},
			patch: func(a args) {
				mockRepo.On("GetLimit", ctx, mockPayload).Return(mockLimit, nil).Once()
				mockRepo.On("GetOngoingPayment", ctx, mockPayload).Return(mockPayment, nil).Once()
				mockRepo.On("GetOngoingPaymentInstallment", ctx, mockInstallmentPayload).Return(mockPaymentInstallment, nil).Once()
			},
		},
		{
			name: "checkLimit_NoError",
			args: args{
				ctx:     ctx,
				payload: mockPayload,
				req:     mockReq,
			},
			expect: expect{
				err: false,
			},
			patch: func(a args) {
				mockRepo.On("GetLimit", ctx, mockPayload).Return(mockEnoughLimit, nil).Once()
				mockRepo.On("GetOngoingPayment", ctx, mockPayload).Return(mockPayment, nil).Once()
				mockRepo.On("GetOngoingPaymentInstallment", ctx, mockInstallmentPayload).Return(mockPaymentInstallment, nil).Once()
			},
		},
	}
	for _, test := range tests {
		svc := getMock()
		test.patch(test.args)
		t.Run(test.name, func(t *testing.T) {
			err := svc.checkLimit(test.args.ctx, test.args.payload, test.args.req)
			if test.expect.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
