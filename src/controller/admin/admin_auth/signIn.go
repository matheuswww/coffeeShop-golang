package admin_auth_controller

import (
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/validation"
	admin_auth_request "matheuswww/coffeeShop-golang/src/controller/model/admin/admin_auth"
	sessionCookie "matheuswww/coffeeShop-golang/src/controller/routes/cookies"
	admin_auth_model "matheuswww/coffeeShop-golang/src/model/admin/admin_auth"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (ac *adminAuthController) SignIn(c *gin.Context) {
	logger.Info("Init SignIn controller", zap.String("journey", "SignIn Controller"))
	var admin_request admin_auth_request.Admin_signIn_request
	if err := c.ShouldBindJSON(&admin_request); err != nil {
		logger.Error("Error trying signIn admin", err, zap.String("journey", "SignIn Controller"))
		rest_err := validation.ValidateError(err)
		c.JSON(rest_err.Code, rest_err)
		return
	}
	domain := admin_auth_model.NewAdminAuthDomainInterface(
		admin_request.Email,
		admin_request.Password,
	)
	err := ac.service.SignIn(domain)
	if err != nil {
		logger.Error("Error trying SignIn", err, zap.String("journey", "SignIn Controller"))
		c.JSON(err.Code, err)
		return
	}
	sessionCookie.SendCoockie(c, int64(domain.GetId()), domain.GetEmail(), "admin")
	c.Status(200)
}
