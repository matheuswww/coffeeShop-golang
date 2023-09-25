package user_auth_service

import (
	"go.uber.org/zap"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	user_auth_model "matheuswww/coffeeShop-golang/src/model/user/user_auth"
	"matheuswww/coffeeShop-golang/src/model/util"
)

func (ud *userAuthDomainService) SignUp(userAuthDomain user_auth_model.UserAuthDomainInterface) *rest_err.RestErr {
	logger.Info("Init SingUp service", zap.String("journey", "SingUp Service"))
	hash, salt, encrypt_err := util.EncryptPassword(userAuthDomain.GetPassword(), nil)
	if encrypt_err != nil {
		logger.Error("Error trying encrypt password", encrypt_err, zap.String("journey", "SingUp Service"))
		return rest_err.NewInternalServerError("server error")
	}
	userAuthDomain.SetEncryptedPassword(hash)
	userAuthDomain.SetSalt(salt)
	err := ud.userRepositroy.SignUp(userAuthDomain)
	if err != nil {
		logger.Error("Error trying create user", err, zap.String("journey", "SingUp Service"))
		return err
	}
	return nil
}
