package admin_auth_controller

import (
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	"matheuswww/coffeeShop-golang/src/configuration/validation"
	admin_auth_request "matheuswww/coffeeShop-golang/src/controller/model/admin/admin_auth"
	sessionCookie "matheuswww/coffeeShop-golang/src/controller/routes/cookies"
	admin_auth_model "matheuswww/coffeeShop-golang/src/model/admin/admin_auth"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (ac *adminAuthController) CreateAdmin(c *gin.Context) {
	logger.Info("Init createAdmin controller", zap.String("journey", "createAdmin Controller"))
	value, coockieErr := sessionCookie.GetCookieValues(c)
	if coockieErr != nil {
		logger.Error("Error invalid coockie", coockieErr, zap.String("journey", "createAdmin Controller"))
		restErr := rest_err.NewBadRequestError(coockieErr.Error())
		c.JSON(restErr.Code, restErr)
		return
	}
	if value.Name != "admin" {
		restErr := rest_err.NewUnauthorizedError("invalid coockie")
		c.JSON(restErr.Code, restErr)
		return
	}
	var user_request admin_auth_request.Admin_signUp_request
	if err := c.ShouldBindJSON(&user_request); err != nil {
		logger.Error("Error trying SingUp user", err, zap.String("journey", "createAdmin Controller"))
		rest_err := validation.ValidateError(err)
		c.JSON(rest_err.Code, rest_err)
		return
	}
	domain := admin_auth_model.NewAdminAuthDomainInterface(
		user_request.Email,
		user_request.Password,
	)
	err := ac.service.CreateAdmin(domain)
	if err != nil {
		logger.Error(
			"Error trying to call createAdmin service",
			err,
			zap.String("journey", "createAdmin Controller"),
		)
		c.JSON(err.Code, err)
		return
	}
	sessionCookie.SendCoockie(c, domain.GetId(), domain.GetEmail(), "admin")
	logger.Info("User created succesfully", zap.String("userId", domain.GetId()), zap.String("journey", "createAdmin Controller"))
	c.Status(http.StatusCreated)
}