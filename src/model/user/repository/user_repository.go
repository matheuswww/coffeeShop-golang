package user_repository

import (
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	user_model "matheuswww/coffeeShop-golang/src/model/user"
)

func NewUserRepository() UserRepository {
	return &userRepository{}
}

type userRepository struct {}

type UserRepository interface {
	CreateUser(user_model.UserDomainInterface) (user_model.UserDomainInterface,*rest_err.RestErr)
}