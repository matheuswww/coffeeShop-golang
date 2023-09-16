package user_controller

import (
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/validation"
	user_request "matheuswww/coffeeShop-golang/src/controller/model/user/request"
	user_model "matheuswww/coffeeShop-golang/src/model/user"
	user_view "matheuswww/coffeeShop-golang/src/view/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) SignUp(c *gin.Context) {
	logger.Info("CreateUser controller",zap.String("journey","CreateUser"))
	var user_request user_request.User_request
	if err := c.ShouldBindJSON(&user_request);err != nil {
		logger.Error("Error trying create user",err,zap.String("journey","createUser"))
		rest_err := validation.ValidateUserError(err)
		c.JSON(rest_err.Code,rest_err)
		return
	}
	domain := user_model.NewUserDomain(
		user_request.Email,
		user_request.Name,
		user_request.Password,
	)
	domainResult,err := uc.service.SignUp(domain)
	if err != nil {
		logger.Error(
			"Error trying to call CreateUser service",
			err,
			zap.String("journey","createUser"),
		)
		c.JSON(err.Code,err)
		return
	}
	IdStr := strconv.FormatInt(domain.GetId(),10)
	logger.Info("User created succesfully",zap.String("userId",IdStr),zap.String("journey","createuser"))
	c.JSON(http.StatusOK,user_view.ConvertDomainToResponse(
		domainResult,
	))
}