package routes

import (
	"database/sql"
	user_routes "matheuswww/coffeeShop-golang/src/controller/routes/user/user_auth_routes"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup,database *sql.DB) {
	user_routes.InitUserAuthRoutes(r,database)
}