package user_profile_service

import (
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud userProfileDomainService) AddToCart(userId string,productId string,quantity int) *rest_err.RestErr {
	logger.Info("Init AddToCart Service", zap.String("journey", "AddToCart Service"))
	err := ud.repository.AddToCart(userId, productId,quantity)
	if err != nil {
		logger.Error("Error trying AddToCart", err, zap.String("journey", "AddToCart Service"))
		return err
	}
	return nil
}
