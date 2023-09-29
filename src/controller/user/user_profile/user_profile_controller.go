package user_profile_controller

import (
	user_profile_service "matheuswww/coffeeShop-golang/src/model/user/user_profile/service"

	"github.com/gin-gonic/gin"
)

func NewUserProfileController(service user_profile_service.UserProfileDomainService) UserProfileController {
	return &userProfileController{
		service,
	}
}

type userProfileController struct {
	service user_profile_service.UserProfileDomainService
}

type UserProfileController interface {
	AddToCart(c *gin.Context)
	GetProfile(c *gin.Context)
}