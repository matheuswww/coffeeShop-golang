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
	databaseConnection *sql.DB
}

type UserAuthRepository interface {
	SignUp(user_auth_model.UserAuthDomainInterface) *rest_err.RestErr
}