package user_auth_controller

import (
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	sessionCookie "matheuswww/coffeeShop-golang/src/controller/routes/cookies"
	user_auth_model "matheuswww/coffeeShop-golang/src/model/user/user_auth"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userAuthControllerInterface) AuthEmail(c *gin.Context) {
	logger.Info("Init AuthEmail Controller", zap.String("journey", "AuthEmail Controller"))
	coockieValues, coockieErr := sessionCookie.GetCookieValues(c)
	if coockieErr != nil {
		logger.Error("Error invalid coockie", coockieErr, zap.String("journey", "AuthEmail Controller"))
		restErr := rest_err.NewBadRequestError(coockieErr.Error())
		c.JSON(restErr.Code, restErr)
		return
	}
	token := c.Param("token")
	userDomain := user_auth_model.NewUserDomainAuthEmail(
		coockieValues.Id,
	)
	err := uc.service.AuthEmail(userDomain, token)
	if err != nil {
		logger.Error("Error trying AuthEmail Controller", err, zap.String("journey", "AuthEmail Controller"))
		c.JSON(err.Code, err)
		return
	}
	c.Status(200)
}
