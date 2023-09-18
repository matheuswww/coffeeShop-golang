package user_auth_controller

import (
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/validation"
	user_request "matheuswww/coffeeShop-golang/src/controller/model/user/request"
	"matheuswww/coffeeShop-golang/src/controller/routes/coockies"
	user_auth_model "matheuswww/coffeeShop-golang/src/model/user/user_auth"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userAuthControllerInterface) SignUp(c *gin.Context) {
	logger.Info("CreateUser controller",zap.String("journey","CreateUser"))
	var user_request user_request.User_request
	if err := c.ShouldBindJSON(&user_request);err != nil {
		logger.Error("Error trying create user",err,zap.String("journey","createUser"))
		rest_err := validation.ValidateUserError(err)
		c.JSON(rest_err.Code,rest_err)
		return
	}
	domain := user_auth_model.NewUserAuthDomain(
		user_request.Email,
		user_request.Name,
		user_request.Password,
	)
	err := uc.service.SignUp(domain)
	if err != nil {
		logger.Error(
			"Error trying to call CreateUser service",
			err,
			zap.String("journey","createUser"),
		)
		c.JSON(err.Code,err)
		return
	}
	coockies.SendCookie(c,domain.GetId())
	logger.Info("User created succesfully",zap.String("userId",strconv.Itoa(domain.GetId())),zap.String("journey","createuser"))
	c.Status(http.StatusOK)
}