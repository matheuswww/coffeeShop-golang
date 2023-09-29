package product_routes

import (
	product_controller "matheuswww/coffeeShop-golang/src/controller/product"
	product_repository "matheuswww/coffeeShop-golang/src/model/product/repository"
	product_service "matheuswww/coffeeShop-golang/src/model/product/service"

	"github.com/gin-gonic/gin"
)

func initAdminProductRoutes(r *gin.RouterGroup) {
	productController := initProductController()
	authGroup := r.Group("/product")

	authGroup.GET("/getAll",productController.GetAll)
}

func initProductController() product_controller.ProductController {
	productRepository := product_repository.NewProductDomainRepository()
	productService := product_service.NewProductService(productRepository)
	productController := product_controller.NewProductController(productService)
	return productController
}