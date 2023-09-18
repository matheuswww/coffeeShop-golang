package user_auth_service

import (
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	user_auth_model "matheuswww/coffeeShop-golang/src/model/user/user_auth"
	user_auth_util "matheuswww/coffeeShop-golang/src/model/user/user_auth/util"
	"go.uber.org/zap"
)

func (ud *userAuthDomainService) SignUp(userAuthDomain user_auth_model.UserAuthDomainInterface) *rest_err.RestErr {
	logger.Info("Init SingUp service",zap.String("journey","SingUp"))
	hash,salt,encrypt_err := user_auth_util.EncryptPass(userAuthDomain.GetPassword())
	if encrypt_err != nil {
		logger.Error("Error trying encrypt password",encrypt_err,zap.String("journey","SingUp"))
		return rest_err.NewInternalServerError("database error")
	}
	userAuthDomain.SetEncryptPassword(hash)
	userAuthDomain.SetSalt(salt)
	err := ud.userRepositroy.SignUp(userAuthDomain)
	if err != nil {
		logger.Error("Error trying create user",err,zap.String("journey","SingUp"))
		return err
	}
	return nil
}