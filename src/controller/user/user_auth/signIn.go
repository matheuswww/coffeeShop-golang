package user_auth_controller

import (
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/validation"
	user_auth_request "matheuswww/coffeeShop-golang/src/controller/model/user/user_auth/request"
	sessionCookie "matheuswww/coffeeShop-golang/src/controller/routes/cookies"

	user_auth_model "matheuswww/coffeeShop-golang/src/model/user/user_auth"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (us *userAuthControllerInterface) SignIn(c *gin.Context) {
	logger.Info("Init SignIn controller", zap.String("journey", "SignIn Controller"))
	var user_request *user_auth_request.User_request_signIn
	if err := c.ShouldBindJSON(&user_request); err != nil {
		logger.Error("Error trying signIn user", err, zap.String("journey", "SignIn Controller"))
		rest_err := validation.ValidateError(err)
		c.JSON(rest_err.Code, rest_err)
		return
	}
	domain := user_auth_model.NewUserSignInDomain(
		user_request.Email,
		user_request.Password,
	)
	err := us.service.SignIn(domain)
	if err != nil {
		logger.Error("Error trying SignIn", err, zap.String("journey", "SignIn Controller"))
		c.JSON(err.Code, err)
		return
	}
	sessionCookie.SendCoockie(c, domain.GetId(), domain.GetEmail(), domain.GetName())
	c.Status(200)
}
