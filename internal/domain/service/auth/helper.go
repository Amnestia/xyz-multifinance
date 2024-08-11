package authsvc

import (
	"context"
	"database/sql"

	"github.com/amnestia/xyz-multifinance/internal/domain/constant"
	authmodel "github.com/amnestia/xyz-multifinance/internal/domain/model/auth"
	"github.com/amnestia/xyz-multifinance/pkg/logger"
)

func (svc *Service) getAccount(ctx context.Context, acc *authmodel.ConsumerAuthRequest) (ret *authmodel.Account, err error) {
	acc.NIK, err = generateHMACHash(acc.NIK, svc.Config.Crypt.HMAC.Key, svc.Config.Crypt.HMAC.Pepper)
	if err != nil {
		err = logger.ErrorWrap(err, "getConsumerData.HMACGetHash")
		return
	}
	ret, err = svc.Repo.Auth(ctx, acc.NIK)
	if err != nil {
		err = logger.ErrorWrap(err, "getConsumerData.RepoAuth")
		if err == sql.ErrNoRows {
			err = constant.LoginFailedError{}
			return
		}
		return
	}
	return
}

func (svc *Service) buildConsumerRegistrationData(ctx context.Context, req *authmodel.RegisterRequest) (acc *authmodel.Account, err error) {
	acc = &authmodel.Account{}
	acc.NIKIndex, err = generateHMACHash(req.NIK, svc.Config.Crypt.HMAC.Key, svc.Config.Crypt.HMAC.Pepper)
	if err != nil {
		err = logger.ErrorWrap(err, "buildConsumerRegistrationData.HMACGetHash")
		return
	}
	acc.NIK, err = encryptAES(acc.NIK, svc.Config.Crypt.AESKey)
	if err != nil {
		err = logger.ErrorWrap(err, "buildConsumerRegistrationData.AESEncryptNIK")
		return
	}
	acc.Fullname, err = encryptAES(acc.Fullname, svc.Config.Crypt.AESKey)
	if err != nil {
		err = logger.ErrorWrap(err, "buildConsumerRegistrationData.AESEncryptNIK")
		return
	}
	acc.LegalName, err = encryptAES(acc.LegalName, svc.Config.Crypt.AESKey)
	if err != nil {
		err = logger.ErrorWrap(err, "buildConsumerRegistrationData.AESEncryptNIK")
		return
	}
	acc.DateOfBirth, err = encryptAES(acc.DateOfBirth, svc.Config.Crypt.AESKey)
	if err != nil {
		err = logger.ErrorWrap(err, "buildConsumerRegistrationData.AESEncryptNIK")
		return
	}
	acc.PlaceOfBirth, err = encryptAES(acc.PlaceOfBirth, svc.Config.Crypt.AESKey)
	if err != nil {
		err = logger.ErrorWrap(err, "buildConsumerRegistrationData.AESEncryptNIK")
		return
	}
	acc.IdentityPhoto, err = encryptAES(acc.IdentityPhoto, svc.Config.Crypt.AESKey)
	if err != nil {
		err = logger.ErrorWrap(err, "buildConsumerRegistrationData.AESEncryptNIK")
		return
	}
	acc.Photo, err = encryptAES(acc.Photo, svc.Config.Crypt.AESKey)
	if err != nil {
		err = logger.ErrorWrap(err, "buildConsumerRegistrationData.AESEncryptNIK")
		return
	}
	acc.Password, err = generateArgonHash(req.Password, svc.Config.Auth.Pepper)
	if err != nil {
		err = logger.ErrorWrap(err, "buildConsumerRegistrationData.GenerateHashPassword")
		return
	}
	acc.PIN, err = generateArgonHash(req.PIN, svc.Config.Auth.Pepper)
	if err != nil {
		err = logger.ErrorWrap(err, "buildConsumerRegistrationData.GenerateHashPIN")
		return
	}

	return
}
