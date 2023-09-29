package user_profile_controller

import (
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/controller/routes/coockies"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userProfileController) GetProfile(c *gin.Context) {
	logger.Info("Init UserProfile controller",zap.String("journey","GetProfile Controller"))
	cookie,cookieErr := coockies.GetCookieValues(c)
	if cookieErr != nil {
		logger.Error("Error invalid coockie",cookieErr,zap.String("journey","GetProfile Controller"))
		c.JSON(401,struct{
			message string
		}{
			"invalid coockie",
		})
	}
	c.JSON(200,struct{
		Id int64 `json:"id"`
		Email string `json:"email"`
		Name string `json:"name"`
	}{
		cookie.Id,
		cookie.Email,
		cookie.Name,
	})
}