package authsvc

import (
	"context"
	"database/sql"

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
		if err == sql.ErrNoRows {
			err = sql.ErrNoRows
			return
		}
		err = logger.ErrorWrap(err, "getConsumerData.RepoAuth")
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
	acc.NIK, err = encryptAES(req.NIK, svc.Config.Crypt.AESKey)
	if err != nil {
		err = logger.ErrorWrap(err, "buildConsumerRegistrationData.AESEncryptNIK")
		return
	}
	acc.Fullname, err = encryptAES(req.Fullname, svc.Config.Crypt.AESKey)
	if err != nil {
		err = logger.ErrorWrap(err, "buildConsumerRegistrationData.AESEncryptFullname")
		return
	}
	acc.LegalName, err = encryptAES(req.LegalName, svc.Config.Crypt.AESKey)
	if err != nil {
		err = logger.ErrorWrap(err, "buildConsumerRegistrationData.AESEncryptLegalName")
		return
	}
	acc.DateOfBirth, err = encryptAES(req.DateOfBirth, svc.Config.Crypt.AESKey)
	if err != nil {
		err = logger.ErrorWrap(err, "buildConsumerRegistrationData.AESEncryptDOB")
		return
	}
	acc.PlaceOfBirth, err = encryptAES(req.PlaceOfBirth, svc.Config.Crypt.AESKey)
	if err != nil {
		err = logger.ErrorWrap(err, "buildConsumerRegistrationData.AESEncryptPlaceOfBirth")
		return
	}
	acc.IdentityPhoto, err = encryptAES(req.IdentityPhoto, svc.Config.Crypt.AESKey)
	if err != nil {
		err = logger.ErrorWrap(err, "buildConsumerRegistrationData.AESEncryptIdentityPhoto")
		return
	}
	acc.Photo, err = encryptAES(req.Photo, svc.Config.Crypt.AESKey)
	if err != nil {
		err = logger.ErrorWrap(err, "buildConsumerRegistrationData.AESEncryptPhoto")
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
