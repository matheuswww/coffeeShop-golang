package admin_auth_controller

import (
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	"matheuswww/coffeeShop-golang/src/configuration/validation"
	admin_auth_request "matheuswww/coffeeShop-golang/src/controller/model/admin/admin_auth"
	sessionCookie "matheuswww/coffeeShop-golang/src/controller/routes/cookies"
	admin_auth_model "matheuswww/coffeeShop-golang/src/model/admin/admin_auth"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (ac *adminAuthController) SignUp(c *gin.Context) {
	logger.Info("Init SignUp controller", zap.String("journey", "SignUp Controller"))
	value, coockieErr := sessionCookie.GetCookieValues(c)
	if coockieErr != nil {
		logger.Error("Error invalid coockie", coockieErr, zap.String("journey", "SignUp Controller"))
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
		logger.Error("Error trying SingUp user", err, zap.String("journey", "SignUp Controller"))
		rest_err := validation.ValidateError(err)
		c.JSON(rest_err.Code, rest_err)
		return
	}
	domain := admin_auth_model.NewAdminAuthDomainInterface(
		user_request.Email,
		user_request.Password,
	)
	err := ac.service.SignUp(domain)
	if err != nil {
		logger.Error(
			"Error trying to call SignUp service",
			err,
			zap.String("journey", "SignUp Controller"),
		)
		c.JSON(err.Code, err)
		return
	}
	sessionCookie.SendCoockie(c, int64(domain.GetId()), domain.GetEmail(), "admin")
	logger.Info("User created succesfully", zap.String("userId", strconv.FormatInt(int64(domain.GetId()), 10)), zap.String("journey", "SignUp Controller"))
	c.Status(http.StatusCreated)
}