package user_repository

import (
	"database/sql"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	user_model "matheuswww/coffeeShop-golang/src/model/user"
)

func NewUserRepository(database *sql.DB) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *sql.DB
}

type UserRepository interface {
	SignUp(user_model.UserDomainInterface,[]byte,[]byte) *rest_err.RestErr
}