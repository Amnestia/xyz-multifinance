package account

import (
	"context"
	"errors"

	authmodel "github.com/amnestia/xyz-multifinance/internal/domain/model/auth"
	"github.com/amnestia/xyz-multifinance/internal/lib/paseto"
	"github.com/amnestia/xyz-multifinance/pkg/logger"
)

// Data struct containing account data from token
type Data struct {
	ID int64
}

// GetData get account data from context inserted from middleware
func GetData(ctx context.Context) (*Data, error) {
	payload := ctx.Value(paseto.AuthData)
	if payload == nil {
		return nil, logger.ErrorWrap(errors.New("nil on context"), "auth", "invalid data")
	}
	p := payload.(paseto.Payload)
	return &Data{
		ID: p.ID,
	}, nil
}

// GetPartnerData get partner data from context inserted from middleware
func GetPartnerData(ctx context.Context) (*authmodel.Partner, error) {
	payload := ctx.Value(paseto.AuthData)
	if payload == nil {
		return nil, logger.ErrorWrap(errors.New("nil on context"), "auth", "invalid data")
	}
	return payload.(*authmodel.Partner), nil
}
