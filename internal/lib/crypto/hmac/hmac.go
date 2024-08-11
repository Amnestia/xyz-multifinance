package libhmac

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
)

func GetHash(target, key, salt string) (ret string, err error) {
	target = base64.StdEncoding.EncodeToString([]byte(target))
	acKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}
	hm := hmac.New(sha512.New, acKey)
	toEncrypt, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		return "", err
	}
	toEncrypt = append(toEncrypt, []byte(target)...)
	_, err = hm.Write([]byte(toEncrypt))
	if err != nil {
		return "", err
	}
	ret = base64.StdEncoding.EncodeToString(hm.Sum(nil))
	return
}
