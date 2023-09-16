package user_service

import (
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	user_model "matheuswww/coffeeShop-golang/src/model/user"
	user_repository "matheuswww/coffeeShop-golang/src/model/user/repository"
)

func NewUserDomainService(userRepository user_repository.UserRepository) UserDomainService{
	return &userDomainService{
		userRepository,
	}
}

type userDomainService struct {
	userRepositroy user_repository.UserRepository
}

type UserDomainService interface {
	SignUp(ud user_model.UserDomainInterface) (user_model.UserDomainInterface,*rest_err.RestErr)
}