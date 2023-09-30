package product_controller

import (
	"context"
	"encoding/json"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
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
	rdb := pc.redis
	cachedData, cacheErr := rdb.Get(ctx,"product:all").Result()
	if cacheErr == nil {
		if err := json.Unmarshal([]byte(cachedData), &products); err != nil {
			logger.Error("Error fetching data from cache",err,zap.String("journey", "GetAll Product Controller"))
		} else {
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