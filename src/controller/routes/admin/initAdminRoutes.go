package admin_routes

import (
	"database/sql"
	admin_auth_routes "matheuswww/coffeeShop-golang/src/controller/routes/admin/admin_auth"
	admin_product_routes "matheuswww/coffeeShop-golang/src/controller/routes/admin/admin_product"

	"github.com/gin-gonic/gin"
)

func InitAdminRoutes(r *gin.RouterGroup, database *sql.DB) {
	admin_auth_routes.InitAdminAuthRoutes(r, database)
	admin_product_routes.InitAdminProductRoutes(r, database)
}
