package user_service

import (
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	user_model "matheuswww/coffeeShop-golang/src/model/user"

	"go.uber.org/zap"
)

func (ud *userDomainService) SignUp(userDomain user_model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init CreateUser service",zap.String("journey","CreateUser"))
	hash,salt,encrypt_err := userDomain.EncryptPassword()
	if encrypt_err != nil {
		logger.Error("Error trying encrypt password",encrypt_err,zap.String("journey","createUser"))
		return rest_err.NewInternalServerError("database error")
	}
	err := ud.userRepositroy.SignUp(userDomain,hash,salt)
	if err != nil {
		logger.Error("Error trying create user",err,zap.String("journey","createUser"))
		return err
	}
	return nil
}