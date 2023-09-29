package product_service

import (
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	product_model "matheuswww/coffeeShop-golang/src/model/product"

	"go.uber.org/zap"
)

func (pd *productDomainService) GetAll(products *[]product_model.ProductDomainInterface) *rest_err.RestErr {
	logger.Info("Init GetAll service", zap.String("journey", "GetAll Service"))
	err := pd.repository.GetAll(products)
	if err != nil {
		logger.Error("Error trying GetAll products", err, zap.String("journey", "GetAll Product"))
		return err
	}
	return nil
}
