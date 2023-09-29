package user_profile_service

import (
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	product_model "matheuswww/coffeeShop-golang/src/model/product"
	user_profile_model "matheuswww/coffeeShop-golang/src/model/user/user_profile"

	"go.uber.org/zap"
)

func (ud userProfileDomainService) AddToCart(userProfileDomainService user_profile_model.UserProfileDomainInterface,productDomain product_model.ProductDomainInterface) *rest_err.RestErr {
	logger.Info("Init AddToCart Service",zap.String("journey","AddToCart Service"))
	err := ud.repository.AddToCart(userProfileDomainService,productDomain)
	if err != nil {
		logger.Error("Error trying AddToCart",err,zap.String("journey","AddToCart Service"))
		return err
	}
	return nil
}