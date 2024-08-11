package authsvc

import (
	"context"

	authmodel "github.com/amnestia/xyz-multifinance/internal/domain/model/auth"
	"github.com/amnestia/xyz-multifinance/internal/lib/crypto/aes"
	"github.com/amnestia/xyz-multifinance/internal/lib/crypto/argon"
	libhmac "github.com/amnestia/xyz-multifinance/internal/lib/crypto/hmac"
)

var (
	getAccount = func(svc *Service, ctx context.Context, acc *authmodel.ConsumerAuthRequest) (ret *authmodel.Account, err error) {
		return svc.getAccount(ctx, acc)
	}

	buildConsumerRegistrationData = func(svc *Service, ctx context.Context, req *authmodel.RegisterRequest) (acc *authmodel.Account, err error) {
		return svc.buildConsumerRegistrationData(ctx, req)
	}
)

var (
	generateArgonHash = func(s string, pepper string) (string, error) {
		return argon.GenerateHash(s, pepper)
	}

	verifyArgonHash = func(s string, p string) (bool, error) {
		return argon.VerifyHash(s, p)
	}

	generateHMACHash = func(s string, key string, pepper string) (string, error) {
		return libhmac.GetHash(s, key, pepper)
	}

	encryptAES = func(s string, key string) (string, error) {
		return aes.Encrypt(s, key)
	}

	decryptAES = func(s string, key string) (string, error) {
		return aes.Decrypt(s, key)
	}
)
