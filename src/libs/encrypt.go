package libs

import (
	"encoding/base64"
	encrypt "libs/encrypt"
)

func Encrypt(str string, key string) (string, error) {
	dstr, err := encrypt.DesEncrypt([]byte(str), []byte(key))
	if err != nil {
		return "", err
	}
	estr := base64.StdEncoding.EncodeToString(dstr)
	return estr, nil
}

func Decrypt(str string, key string) (string, error) {
	estr, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	dstr, err := encrypt.DesDecrypt([]byte(estr), []byte(key))
	if err != nil {
		return "", err
	}
	return string(dstr), nil
}
