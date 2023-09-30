package product_repository

import (
	"context"
	"encoding/json"
	"errors"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	user_profile_response "matheuswww/coffeeShop-golang/src/controller/model/user/user_profile/response"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func (pr *productRepository) GetAll(rdb *redis.Client,ctxRedis *context.Context) ([]user_profile_response.Product, *rest_err.RestErr) {
	logger.Info("Init GetAll repository", zap.String("journey", "GetAll Repository"))
	db := pr.database
	ctx, cancel := context.WithTimeout(context.Background(), (time.Second * 5))
	defer cancel()
	query := "SELECT uuid,name,price,stock FROM products"
	result, err := db.QueryContext(ctx, query)
	if err != nil {
		logger.Error("Error trying to GetAll products", err, zap.String("journey", "GetAll repository"))
		return nil, rest_err.NewInternalServerError("server error")
	}
	defer result.Close()
	var JSONlist []user_profile_response.Product
	for result.Next() {
		var id, name string
		var price float32
		var stock int
		if err := result.Scan(&id, &name, &price, &stock); err != nil {
			logger.Error("Error scanning result", err, zap.String("journey", "GetAll Repository"))
			return nil, rest_err.NewInternalServerError("server error")
		}
		product := user_profile_response.Product{
			ID:    id,
			Name:  name,
			Price: price,
			Stock: stock,
		}
		JSONlist = append(JSONlist, product)
	}
	if len(JSONlist) == 0 {
		logger.Error("Error no products found", errors.New("no products found"), zap.String("journey", "GetAll Repository"))
		return nil, rest_err.NewNotFoundError("no products")
	}
	b,err := json.Marshal(JSONlist)
	if err != nil {
		logger.Error("Error trying marshal jsonlist",err,zap.String("journey","GetAll repository"))
		return nil,rest_err.NewInternalServerError("server error")
	}
	cacheErr := rdb.Set(*ctxRedis, "product:all", b, time.Hour)
	if cacheErr.Err() != nil {
		logger.Error("Error trying set cache",cacheErr.Err(),zap.String("journey","GetAll Repository"))
	}
	return JSONlist, nil
}