package user_profile_service

import (
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
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
	AddToCart(userId string,productId string,quantity int) *rest_err.RestErr
}
