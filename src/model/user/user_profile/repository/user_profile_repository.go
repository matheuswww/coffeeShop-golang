package user_profile_repository

import (
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	product_model "matheuswww/coffeeShop-golang/src/model/product"
	user_profile_model "matheuswww/coffeeShop-golang/src/model/user/user_profile"
)

func NewUserProfileRepository() UserProfileRepository {
	return &userProfileRepository{}
}

type userProfileRepository struct{}

type UserProfileRepository interface {
	AddToCart(userProfileDomainSevice user_profile_model.UserProfileDomainInterface, productDomain product_model.ProductDomainInterface) *rest_err.RestErr
}
