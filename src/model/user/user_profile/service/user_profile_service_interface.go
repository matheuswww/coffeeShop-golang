package user_profile_service

import (
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	product_model "matheuswww/coffeeShop-golang/src/model/product"
	user_profile_model "matheuswww/coffeeShop-golang/src/model/user/user_profile"
	user_profile_repository "matheuswww/coffeeShop-golang/src/model/user/user_profile/repository"
)

func NewUserProfileService(repository user_profile_repository.UserProfileRepository) UserProfileDomainService {
	return &userProfileDomainService{
		repository,
	}
}

type userProfileDomainService struct {
	repository user_profile_repository.UserProfileRepository
}

type UserProfileDomainService interface {
	AddToCart(userProfileDomainService user_profile_model.UserProfileDomainInterface, productDomain product_model.ProductDomainInterface) *rest_err.RestErr
}
