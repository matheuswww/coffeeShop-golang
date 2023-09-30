package routes

import (
	"database/sql"
	admin_routes "matheuswww/coffeeShop-golang/src/controller/routes/admin"
	product_routes "matheuswww/coffeeShop-golang/src/controller/routes/product"
	user_routes "matheuswww/coffeeShop-golang/src/controller/routes/user"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, database *sql.DB) {
	user_routes.InitUserRoutes(r, database)
	admin_routes.InitAdminRoutes(r, database)
	product_routes.InitProductRoutes(r, database)
}
