package user_routes

import (
	"matheuswww/coffeeShop-golang/src/controller/routes/user/user_auth_routes"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.RouterGroup) {
	user_auth_routes.InitUserAuthRoutes(r)
}