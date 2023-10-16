package user_profile_repository

import (
	"database/sql"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
)

func NewUserProfileRepository(database *sql.DB) UserProfileRepository {
	return &userProfileRepository{
		database,
	}
}

type userProfileRepository struct {
	database *sql.DB
}

type UserProfileRepository interface {
	AddToCart(userId string, productId string,quantity int) *rest_err.RestErr
}
