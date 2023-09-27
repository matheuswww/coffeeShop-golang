package admin_auth_controller

import (
	admin_auth_service "matheuswww/coffeeShop-golang/src/model/admin/admin_auth/service"

	"github.com/gin-gonic/gin"
)

func NewAdminAuthInterface(AdminAuthService admin_auth_service.AdminAuthDomainService) AdminAuthControllerInterface {
	return &adminAuthController{
		AdminAuthService,
	}
}

type AdminAuthControllerInterface interface {
	SignIn(c *gin.Context)
}

type adminAuthController struct {
	service admin_auth_service.AdminAuthDomainService
}
