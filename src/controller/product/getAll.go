package product_controller

import (
	"context"
	"encoding/json"
	"fmt"

	"matheuswww/coffeeShop-golang/src/configuration/logger"
	redisClient "matheuswww/coffeeShop-golang/src/configuration/redis"
	user_profile_response "matheuswww/coffeeShop-golang/src/controller/model/user/user_profile/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (pc *productController) GetAll(c *gin.Context) {
	logger.Info("Init GetAll controller", zap.String("journey", "GetAll Controller"))
	var products []user_profile_response.Product
	ctx,cancel := context.WithTimeout(context.Background(),time.Second * 10)
	defer cancel()
	rdb,redisErr := redisClient.NewRedis().NewRedisConnection()
	if redisErr != nil {
		logger.Error("Error trying connect redis",redisErr,zap.String("journey","GetAll"))
	}
	defer rdb.Close()
	cachedData, cacheErr := rdb.Get(ctx,"product:all").Result()
	if cacheErr == nil {
		if err := json.Unmarshal([]byte(cachedData), &products); err != nil {
			logger.Error("Error fetching data from cache",err,zap.String("journey", "GetAll Product Controller"))
		} else {
			fmt.Println("cache")
			c.JSON(http.StatusOK,products)
			return
		}
	}
	products,err := pc.service.GetAll(rdb,&ctx)
	if err != nil {
		logger.Error("Error trying GetAll products",err,zap.String("journey","GetAll Controller"))
		c.JSON(err.Code,err)
		return
	}
	c.JSON(http.StatusOK, products)
}