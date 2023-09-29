package user_auth_repository

import (
	"database/sql"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	user_auth_model "matheuswww/coffeeShop-golang/src/model/user/user_auth"
)

func NewUserAuthRepository(database *sql.DB) UserAuthRepository {
	return &userAuthRepository{
		database,
	}
}

type userAuthRepository struct {
	database *sql.DB
}

type UserAuthRepository interface {
	SignUp(user_auth_model.UserAuthDomainInterface) *rest_err.RestErr
	SignIn(userDomain user_auth_model.UserAuthDomainInterface) *rest_err.RestErr
	SendAuthEmail(userDomain user_auth_model.UserAuthDomainInterface, token string) *rest_err.RestErr
	AuthEmail(userDomain user_auth_model.UserAuthDomainInterface) *rest_err.RestErr
}
