package routes

import (
	"database/sql"
	"matheuswww/coffeeShop-golang/src/controller/routes/user_routes"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup,database *sql.DB) {
	user_routes.InitUserRoutes(r,database)
}