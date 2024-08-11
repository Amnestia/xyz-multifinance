package auth

import (
	"fmt"
	"strings"

	authmodel "github.com/amnestia/xyz-multifinance/internal/domain/model/auth"
)

func validateRegister(req *authmodel.RegisterRequest) (err error) {
	req.NIK = strings.TrimSpace(req.NIK)
	req.Password = strings.TrimSpace(req.Password)
	req.PIN = strings.TrimSpace(req.PIN)
	req.Fullname = strings.TrimSpace(req.Fullname)
	req.LegalName = strings.TrimSpace(req.LegalName)
	req.DateOfBirth = strings.TrimSpace(req.DateOfBirth)
	req.PlaceOfBirth = strings.TrimSpace(req.PlaceOfBirth)
	req.IdentityPhoto = strings.TrimSpace(req.IdentityPhoto)
	req.Photo = strings.TrimSpace(req.Photo)
	if req.NIK == "" {
		return fmt.Errorf("NIK is required")
	}
	if req.Fullname == "" {
		return fmt.Errorf("Fullname is required")
	}
	if req.LegalName == "" {
		return fmt.Errorf("Legal Name is required")
	}
	if req.DateOfBirth == "" {
		return fmt.Errorf("Legal Name is required")
	}
	if req.PlaceOfBirth == "" {
		return fmt.Errorf("Legal Name is required")
	}
	if req.IdentityPhoto == "" {
		return fmt.Errorf("Legal Name is required")
	}
	if req.Photo == "" {
		return fmt.Errorf("Legal Name is required")
	}
	if req.Salary < 1 {
		return fmt.Errorf("Invalid salary")
	}
	if err = validatePassword(req.Password); err != nil {
		return
	}
	if err = validatePIN(req.PIN); err != nil {
		return
	}
	return
}

func validatePIN(pin string) (err error) {
	if len(pin) != 6 {
		return fmt.Errorf("Invalid PIN format")
	}
	return
}

func validatePassword(password string) (err error) {
	if len(password) < 16 {
		return fmt.Errorf("Password should be at least 16 characters")
	}
	if len(password) > 256 {
		return fmt.Errorf("Password should not be more than 256 characters")
	}
	return
}
