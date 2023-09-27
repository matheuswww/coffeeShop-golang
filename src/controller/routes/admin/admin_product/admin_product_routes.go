package admin_product_routes

import (
	admin_product_controller "matheuswww/coffeeShop-golang/src/controller/admin/admin_product"
	"matheuswww/coffeeShop-golang/src/controller/routes/coockies"
	admin_product_repository "matheuswww/coffeeShop-golang/src/model/admin/admin_product/repository"
	admin_product_service "matheuswww/coffeeShop-golang/src/model/admin/admin_product/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InitAdminProductRoutes(r *gin.RouterGroup) {
	userController := initAdminProductController()
	authGroup := r.Group("/admin/product")
	authGroup.Use(sessions.Sessions("auth", coockies.Store()))

	authGroup.POST("/insert",userController.InsertProduct)
}

func initAdminProductController() admin_product_controller.AdminProductController {
	adminProductRepository := admin_product_repository.NewAdminProductRepository()
	adminProductService := admin_product_service.NewAdminProductService(adminProductRepository)
	adminProductController := admin_product_controller.NewAdminProductController(adminProductService)
	return adminProductController
}