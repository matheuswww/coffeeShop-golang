package user_profile_controller

import (
	"fmt"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/validation"
	user_profile_request "matheuswww/coffeeShop-golang/src/controller/model/user/user_profile/request"
	sessionCookie "matheuswww/coffeeShop-golang/src/controller/routes/cookies"
	product_model "matheuswww/coffeeShop-golang/src/model/product"
	user_profile_model "matheuswww/coffeeShop-golang/src/model/user/user_profile"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userProfileController) AddToCart(c *gin.Context) {
	logger.Info("Init UserProfile controller", zap.String("journey", "AddToCart Controller"))
	cookie, cookieErr := sessionCookie.GetCookieValues(c)
	if cookieErr != nil {
		logger.Error("Error invalid coockie", cookieErr, zap.String("journey", "AddToCart Controller"))
		c.JSON(401, struct {
			message string
		}{
			"invalid coockie",
		})
	}
	var userProfileRequest user_profile_request.User_profile_request_addToCart
	if err := c.ShouldBindJSON(&userProfileRequest); err != nil {
		restErr := validation.ValidateError(err)
		c.JSON(restErr.Code, restErr)
		return
	}
	userProfileDomain := user_profile_model.NewUserProfileDomain(
		cookie.Id,
		cookie.Email,
		cookie.Name,
	)
	fmt.Println(userProfileRequest.ProductId)
	productDomain := product_model.NewProductDomainService(
		userProfileRequest.ProductId,
		userProfileRequest.ProductName,
		userProfileRequest.ProductPrice,
		userProfileRequest.ProductQuantity,
	)
	err := uc.service.AddToCart(userProfileDomain, productDomain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.Status(201)
}
