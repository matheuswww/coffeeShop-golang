package user_auth_controller

import (
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	sessionCookie "matheuswww/coffeeShop-golang/src/controller/routes/cookies"
	user_auth_model "matheuswww/coffeeShop-golang/src/model/user/user_auth"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userAuthControllerInterface) SendAuthEmail(c *gin.Context) {
	logger.Info("Init SendAuthEmail controller", zap.String("journey", "SendAuthEmail Controller"))
	coockieValues, coockieErr := sessionCookie.GetCookieValues(c)
	if coockieErr != nil {
		logger.Error("Error invalid coockie", coockieErr, zap.String("journey", "AuthEmail Controller"))
		restErr := rest_err.NewBadRequestError(coockieErr.Error())
		c.JSON(restErr.Code, restErr)
		return
	}
	userDomain := user_auth_model.NewUserDomainSendAuthEmail(
		coockieValues.Id,
		coockieValues.Email,
		coockieValues.Name,
	)
	err := uc.service.SendAuthEmail(userDomain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.Status(200)
}
