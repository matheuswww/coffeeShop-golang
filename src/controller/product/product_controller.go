package product_controller

import (
	product_service "matheuswww/coffeeShop-golang/src/model/product/service"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func NewProductController(service product_service.ProductDomainRepository,redis *redis.Client) ProductController {
	return &productController{
		service,
		redis,
	}
}

type productController struct {
	service product_service.ProductDomainRepository
	redis *redis.Client
}

type ProductController interface {
	GetAll(c *gin.Context)
}
