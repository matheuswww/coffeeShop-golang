package admin_product_routes

import (
	"database/sql"
	admin_product_controller "matheuswww/coffeeShop-golang/src/controller/admin/admin_product"
	sessionCookie "matheuswww/coffeeShop-golang/src/controller/routes/cookies"
	admin_product_repository "matheuswww/coffeeShop-golang/src/model/admin/admin_product/repository"
	admin_product_service "matheuswww/coffeeShop-golang/src/model/admin/admin_product/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InitAdminProductRoutes(r *gin.RouterGroup, database *sql.DB) {
	userController := initAdminProductController(database)
	authGroup := r.Group("/admin/product")
	authGroup.Use(sessions.Sessions("auth", sessionCookie.Store()))

	authGroup.POST("/insert", userController.InsertProduct)
}

func initAdminProductController(database *sql.DB) admin_product_controller.AdminProductController {
	adminProductRepository := admin_product_repository.NewAdminProductRepository(database)
	adminProductService := admin_product_service.NewAdminProductService(adminProductRepository)
	adminProductController := admin_product_controller.NewAdminProductController(adminProductService)
	return adminProductController
}
