package product_service

import (
	"context"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	user_profile_response "matheuswww/coffeeShop-golang/src/controller/model/user/user_profile/response"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func (pd *productDomainService) GetAll(rdb *redis.Client,ctxRedis *context.Context) ([]user_profile_response.Product,*rest_err.RestErr) {
	logger.Info("Init GetAll service", zap.String("journey", "GetAll Service"))
	products,err := pd.repository.GetAll(rdb,ctxRedis)
	if err != nil {
		logger.Error("Error trying GetAll products", err, zap.String("journey", "GetAll Product"))
		return nil,err
	}
	return products,nil
}
