package product_routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitAdminRoutes(r *gin.RouterGroup, database *sql.DB) {
	initAdminProductRoutes(r, database)
}
