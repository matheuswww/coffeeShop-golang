package admin_auth_service

import (
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	admin_auth_model "matheuswww/coffeeShop-golang/src/model/admin/admin_auth"

	"go.uber.org/zap"
)

func (ad *adminAuthDomainService) SignIn(admin admin_auth_model.AdminAuthDomainInterface) *rest_err.RestErr {
	logger.Info("Init SignIn service", zap.String("journey", "SignIn Service"))
	err := ad.repository.SignIn(admin)
	if err != nil {
		logger.Error("Error trying SingIn", err, zap.String("journey", "SignIn Service"))
		return err
	}
	return nil
}
