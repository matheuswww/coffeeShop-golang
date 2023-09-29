package user_routes

import (
	"database/sql"
	"matheuswww/coffeeShop-golang/src/controller/routes/user/user_auth_routes"
	"matheuswww/coffeeShop-golang/src/controller/routes/user/user_profile_routes"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.RouterGroup, database *sql.DB) {
	user_auth_routes.InitUserAuthRoutes(r, database)
	user_profile_routes.InitUserProfileRoutes(r, database)
}
