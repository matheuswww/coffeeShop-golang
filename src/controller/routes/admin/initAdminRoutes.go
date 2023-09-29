package admin_routes

import (
	"github.com/gin-gonic/gin"
	admin_auth_routes "matheuswww/coffeeShop-golang/src/controller/routes/admin/admin_auth"
	admin_product_routes "matheuswww/coffeeShop-golang/src/controller/routes/admin/admin_product"
)

func InitAdminRoutes(r *gin.RouterGroup) {
	admin_auth_routes.InitAdminAuthRoutes(r)
	admin_product_routes.InitAdminProductRoutes(r)
}
