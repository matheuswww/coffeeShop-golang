package user_auth_service

import (
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	user_auth_model "matheuswww/coffeeShop-golang/src/model/user/user_auth"
	user_auth_repository "matheuswww/coffeeShop-golang/src/model/user/user_auth/repository"
)

func NewUserAuthDomainService(userAuthRepository user_auth_repository.UserAuthRepository) UserAuthDomainService{
	return &userAuthDomainService{
		userAuthRepository,
	}
}

type userAuthDomainService struct {
	userRepositroy user_auth_repository.UserAuthRepository
}

type UserAuthDomainService interface {
	SignUp(ud user_auth_model.UserAuthDomainInterface) *rest_err.RestErr
	SignIn(ud user_auth_model.UserAuthDomainInterface) *rest_err.RestErr
}