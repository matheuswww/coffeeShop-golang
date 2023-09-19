package routes

import (
	"database/sql"
	user_routes "matheuswww/coffeeShop-golang/src/controller/routes/user/user_auth_routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup,database *sql.DB) {
	r.Use(Headers)
	user_routes.InitUserAuthRoutes(r,database)
}

func Headers(c *gin.Context) {
	c.Header("X-Permitted-Cross-Domain-Policies", "none")
	if c.Request.Method == "TRACE" {
		c.JSON(http.StatusMethodNotAllowed,gin.H{
			"message": "TRACE method not allowed",
		})
		c.Abort()
		return
	}
	c.Next()
} 
