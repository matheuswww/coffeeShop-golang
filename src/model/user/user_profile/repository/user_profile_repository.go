package user_profile_repository

import (
	"database/sql"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	product_model "matheuswww/coffeeShop-golang/src/model/product"
	user_profile_model "matheuswww/coffeeShop-golang/src/model/user/user_profile"
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
	AddToCart(userProfileDomainSevice user_profile_model.UserProfileDomainInterface, productDomain product_model.ProductDomainInterface) *rest_err.RestErr
}
