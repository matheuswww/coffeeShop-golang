package admin_product_controller

import (
	admin_product_service "matheuswww/coffeeShop-golang/src/model/admin/admin_product/service"

	"github.com/gin-gonic/gin"
)

func NewAdminProductController(adminProductService admin_product_service.AdminProductDomainService) AdminProductController {
	return &adminProductController{
		adminProductService,
	}
}

type adminProductController struct {
	service admin_product_service.AdminProductDomainService
}

type AdminProductController interface {
	InsertProduct(c *gin.Context)
}
