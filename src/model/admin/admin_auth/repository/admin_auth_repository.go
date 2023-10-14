package admin_auth_repository

import (
	"database/sql"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	admin_auth_model "matheuswww/coffeeShop-golang/src/model/admin/admin_auth"
)

func NewAdminAuthRepository(database *sql.DB) AdminAuthRepository {
	return &adminAuthRepository{
		database,
	}
}

type adminAuthRepository struct {
	database *sql.DB
}

type AdminAuthRepository interface {
	SignIn(admin admin_auth_model.AdminAuthDomainInterface) *rest_err.RestErr
	SignUp(adminDomain admin_auth_model.AdminAuthDomainInterface) *rest_err.RestErr 
}
