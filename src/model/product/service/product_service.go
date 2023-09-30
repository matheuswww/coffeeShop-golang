package product_service

import (
	"context"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	user_profile_response "matheuswww/coffeeShop-golang/src/controller/model/user/user_profile/response"
	product_repository "matheuswww/coffeeShop-golang/src/model/product/repository"

	"github.com/redis/go-redis/v9"
)

func NewProductService(repository product_repository.ProductRepository) ProductDomainRepository {
	return &productDomainService{
		repository: repository,
	}
}

type productDomainService struct {
	repository product_repository.ProductRepository
}

type ProductDomainRepository interface {
	GetAll(rdb *redis.Client,ctxRedis *context.Context) ([]user_profile_response.Product,*rest_err.RestErr)
}
