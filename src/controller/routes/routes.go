package routes

import (
	user_controller "matheuswww/coffeeShop-golang/src/controller/user"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.RouterGroup,userController user_controller.UserControllerInterface) {
	r.POST("/crateUser",userController.CreateUser)
}