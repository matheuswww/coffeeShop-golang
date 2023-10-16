package user_auth_service

import (
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	user_auth_model "matheuswww/coffeeShop-golang/src/model/user/user_auth"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (ur *userAuthDomainService) SignIn(UserAuthDomain user_auth_model.UserAuthDomainInterface) *rest_err.RestErr {
	logger.Info("Init SignIn service", zap.String("journey", "SignIn Service"))
	id,idErr := uuid.NewUUID()
	if idErr != nil {
		logger.Error("Error generating uuid",idErr,zap.String("journey","SignIn service"))
		return rest_err.NewInternalServerError("server error")
	}
	UserAuthDomain.SetId(id.String())
	err := ur.userRepositroy.SignIn(UserAuthDomain)
	if err != nil {
		logger.Error("Error trying SingIn", err, zap.String("journey", "SignIn Service"))
		return err
	}
	return nil
}
