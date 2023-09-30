package product_routes

import (
	"database/sql"
	product_controller "matheuswww/coffeeShop-golang/src/controller/product"
	product_repository "matheuswww/coffeeShop-golang/src/model/product/repository"
	product_service "matheuswww/coffeeShop-golang/src/model/product/service"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func initProductRoutes(r *gin.RouterGroup, database *sql.DB,redis *redis.Client) {
	productController := initProductController(database,redis)
	authGroup := r.Group("/product")

	authGroup.GET("/getAll", productController.GetAll)
}

func initProductController(database *sql.DB,redis *redis.Client) product_controller.ProductController {
	productRepository := product_repository.NewProductDomainRepository(database,redis)
	productService := product_service.NewProductService(productRepository)
	productController := product_controller.NewProductController(productService,redis)
	return productController
}
