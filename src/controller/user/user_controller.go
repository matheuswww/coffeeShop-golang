package user_controller

import (
	user_service "matheuswww/coffeeShop-golang/src/model/user/service"

	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(serviceInterface user_service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		serviceInterface,
	}
}

type userControllerInterface struct {
	service user_service.UserDomainService
}

type UserControllerInterface interface {
	SignUp(c *gin.Context)
}