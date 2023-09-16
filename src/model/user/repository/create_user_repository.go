package user_repository

import (
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	user_model "matheuswww/coffeeShop-golang/src/model/user"

	"go.uber.org/zap"
)

func (ur userRepository) CreateUser(userDomain user_model.UserDomainInterface) (user_model.UserDomainInterface,*rest_err.RestErr){
	logger.Info("Init createUser repository",zap.String("journey","CreateUser"))
	return userDomain,nil
}