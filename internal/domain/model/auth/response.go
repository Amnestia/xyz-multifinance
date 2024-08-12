package authmodel

import "github.com/amnestia/xyz-multifinance/internal/domain/model/common"

// LoginResponse response returned on login
type LoginResponse struct {
	TokenData
	common.DefaultResponse
}

// NewPartnerResponse response returned on creating a new partner
type NewPartnerResponse struct {
	ClientID string
	APIKey   string
	common.DefaultResponse
}
