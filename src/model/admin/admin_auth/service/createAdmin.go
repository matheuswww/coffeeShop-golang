package admin_auth_service

import (
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	admin_auth_model "matheuswww/coffeeShop-golang/src/model/admin/admin_auth"
	"matheuswww/coffeeShop-golang/src/model/util"

	"go.uber.org/zap"
)

func (as *adminAuthDomainService) CreateAdmin(adminDomain admin_auth_model.AdminAuthDomainInterface) *rest_err.RestErr {
	logger.Info("Init CreateAdmin service", zap.String("journey", "CreateAdmin Service"))
	hash, salt, encrypt_err := util.EncryptPassword(adminDomain.GetPassword(),nil)
	if encrypt_err != nil {
		logger.Error("Error trying encrypt password", encrypt_err, zap.String("journey", "CreateAdmin Service"))
		return rest_err.NewInternalServerError("server error")
	}
	adminDomain.SetEncryptedPassword(hash)
	adminDomain.SetSalt(salt)
	err := as.repository.CreateAdmin(adminDomain)
	if err != nil {
		logger.Error("Error trying create user", err, zap.String("journey", "CreateAdmin Service"))
		return err
	}
	return nil
}