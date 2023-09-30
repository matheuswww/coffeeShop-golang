package product_routes

import (
	"database/sql"
	"log"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	redisClient "matheuswww/coffeeShop-golang/src/configuration/redis"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func InitProductRoutes(r *gin.RouterGroup, database *sql.DB) {
	initProductRoutes(r, database,loadRedis())
}

func loadRedis() *redis.Client {
	rdb,redisErr := redisClient.NewRedis().NewRedisConnection()
	if redisErr != nil {
		logger.Error("Error trying connect redis",redisErr,zap.String("journey","loadRedis"))
		log.Fatal("Error trying connect redis")
	}
	return rdb
}
