package product_controller

import (
	product_service "matheuswww/coffeeShop-golang/src/model/product/service"

	"github.com/gin-gonic/gin"
)

func NewProductController(service product_service.ProductDomainRepository) ProductController {
	return &productController{
		service: service,
	}
}

type productController struct {
	service product_service.ProductDomainRepository
}

type ProductController interface {
	GetAll(c *gin.Context)
}
