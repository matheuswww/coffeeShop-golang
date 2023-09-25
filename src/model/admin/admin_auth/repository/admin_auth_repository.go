package admin_auth_repository

import (
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	admin_auth_model "matheuswww/coffeeShop-golang/src/model/admin/admin_auth"
)

func NewAdminAuthRepository() AdminAuthRepository {
	return &adminAuthRepository{}
}

type adminAuthRepository struct{}

type AdminAuthRepository interface {
	SignIn(admin admin_auth_model.AdminAuthDomainInterface) *rest_err.RestErr
}
