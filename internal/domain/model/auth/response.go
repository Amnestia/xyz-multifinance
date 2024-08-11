package authmodel

import "github.com/amnestia/xyz-multifinance/internal/domain/model/common"

// LoginResponse response returned on login
type LoginResponse struct {
	TokenData
	common.DefaultResponse
}
