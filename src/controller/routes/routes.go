package routes

import (
	admin_routes "matheuswww/coffeeShop-golang/src/controller/routes/admin"
	product_routes "matheuswww/coffeeShop-golang/src/controller/routes/product"
	user_routes "matheuswww/coffeeShop-golang/src/controller/routes/user"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	user_routes.InitUserRoutes(r)
	admin_routes.InitAdminRoutes(r)
	product_routes.InitAdminRoutes(r)
}
