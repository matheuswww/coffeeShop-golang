package user_controller

import (
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/validation"
	"matheuswww/coffeeShop-golang/src/controller/model/user_request"
	user_model "matheuswww/coffeeShop-golang/src/model/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("CreateUser controller",zap.String("journey","CreateUser"))
	var user_request user_request.User_request
	if err := c.ShouldBindJSON(&user_request);err != nil {
		logger.Error("CreateUser controller",err)
		rest_err := validation.ValidateUserError(err)
		c.JSON(rest_err.Code,rest_err)
		return
	}
	user := user_model.NewUserDomain("joaodocapa@gmail.com","joaodocapa","joazindocapa")
	uc.service.CreateUser(user)
}