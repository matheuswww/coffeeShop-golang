package user_auth_controller

import (
	user_auth_service "matheuswww/coffeeShop-golang/src/model/user/user_auth/service"

	"github.com/gin-gonic/gin"
)

func NewUserAuthControllerInterface(serviceInterface user_auth_service.UserAuthDomainService) UserAuthControllerInterface {
	return &userAuthControllerInterface{
		serviceInterface,
	}
}

type userAuthControllerInterface struct {
	service user_auth_service.UserAuthDomainService
}

type UserAuthControllerInterface interface {
	SignUp(c *gin.Context)
}