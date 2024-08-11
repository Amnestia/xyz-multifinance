package auth

import (
	"fmt"

	authmodel "github.com/amnestia/xyz-multifinance/internal/domain/model/auth"
)

func validateRegister(req *authmodel.Account) (err error) {
	if req.Username == "" {
		return fmt.Errorf("Username is required")
	}
	if req.Email == "" {
		return fmt.Errorf("Email is required")
	}
	if err = validatePassword(req.Password); err != nil {
		return
	}
	return
}

func validatePassword(password string) (err error) {
	if len(password) < 8 {
		return fmt.Errorf("Password should be at least 8 characters")
	}
	return
}
