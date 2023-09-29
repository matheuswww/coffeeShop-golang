package product_routes

import (
	"database/sql"
	product_controller "matheuswww/coffeeShop-golang/src/controller/product"
	product_repository "matheuswww/coffeeShop-golang/src/model/product/repository"
	product_service "matheuswww/coffeeShop-golang/src/model/product/service"

	"github.com/gin-gonic/gin"
)

func initAdminProductRoutes(r *gin.RouterGroup, database *sql.DB) {
	productController := initProductController(database)
	authGroup := r.Group("/product")

	authGroup.GET("/getAll", productController.GetAll)
}

func initProductController(database *sql.DB) product_controller.ProductController {
	productRepository := product_repository.NewProductDomainRepository(database)
	productService := product_service.NewProductService(productRepository)
	productController := product_controller.NewProductController(productService)
	return productController
}
