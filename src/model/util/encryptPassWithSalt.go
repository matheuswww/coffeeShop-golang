package util

import (
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
)

func EncryptUserPasswordWithSalt(password string, salt []byte, callback func(hash, salt []byte)) *rest_err.RestErr {
	hash, _, err := EncryptPassword(password, salt)
	if err != nil {

		return rest_err.NewInternalServerError("database error")
	}
	callback(hash, salt)
	return nil
}
