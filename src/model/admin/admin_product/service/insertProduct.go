package admin_product_service

import (
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	admin_product_model "matheuswww/coffeeShop-golang/src/model/admin/admin_product"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (as *adminProductDomainService) InsertProduct(adminProductDomain admin_product_model.AdminProductDomainInterface) *rest_err.RestErr {
	logger.Info("Init InsertProduct Service",zap.String("journey","InsertProduct Service"))
	uuidErr := as.generateUUID(adminProductDomain)
	if uuidErr != nil {
		logger.Error("Error trying create UUID",uuidErr,zap.String("journey","InsertProduct Service"))
		return rest_err.NewInternalServerError("server error")
	}
	err := as.repository.InsertProduct(adminProductDomain)
	if err != nil {
		logger.Error("Error trying InsertProduct Service",err,zap.String("journey","InsertProduct Service"))
		return err
	}
	return nil
}

func (as *adminProductDomainService) generateUUID(AdminProductDomain admin_product_model.AdminProductDomainInterface) error {
	u, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	AdminProductDomain.SetUUID(u.String())
	return nil
}