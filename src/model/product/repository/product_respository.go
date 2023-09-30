package product_repository

import (
	"context"
	"database/sql"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	user_profile_response "matheuswww/coffeeShop-golang/src/controller/model/user/user_profile/response"

	"github.com/redis/go-redis/v9"
)

func NewProductDomainRepository(database *sql.DB,redis *redis.Client) ProductRepository {
	return &productRepository{
		database,
	}
}

type productRepository struct {
	database *sql.DB
}

type ProductRepository interface {
	GetAll(rdb *redis.Client,ctxRedis *context.Context) ([]user_profile_response.Product,*rest_err.RestErr)
}
