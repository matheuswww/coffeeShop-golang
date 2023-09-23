package user_auth_repository

import (
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	user_auth_model "matheuswww/coffeeShop-golang/src/model/user/user_auth"
)

func NewUserAuthRepository() UserAuthRepository {
	return &userAuthRepository{}
}

type userAuthRepository struct{}

type UserAuthRepository interface {
	SignUp(user_auth_model.UserAuthDomainInterface) *rest_err.RestErr
	SignIn(userDomain user_auth_model.UserAuthDomainInterface) *rest_err.RestErr
	SendAuthEmail(userDomain user_auth_model.UserAuthDomainInterface, token string) *rest_err.RestErr
	AuthEmail(userDomain user_auth_model.UserAuthDomainInterface) *rest_err.RestErr
}
