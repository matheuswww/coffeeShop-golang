package product_routes

import "github.com/gin-gonic/gin"

func InitAdminRoutes(r *gin.RouterGroup) {
	initAdminProductRoutes(r)
}
