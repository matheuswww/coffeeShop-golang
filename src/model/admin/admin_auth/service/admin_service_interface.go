package admin_auth_service

import (
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	admin_auth_model "matheuswww/coffeeShop-golang/src/model/admin/admin_auth"
	admin_auth_repository "matheuswww/coffeeShop-golang/src/model/admin/admin_auth/repository"
)

func NewAuthDomainService(AdminAuthRepository admin_auth_repository.AdminAuthRepository) AdminAuthDomainService {
	return &adminAuthDomainService{
		AdminAuthRepository,
	}
}

type adminAuthDomainService struct {
	repository admin_auth_repository.AdminAuthRepository
}

type AdminAuthDomainService interface {
	SignIn(admin admin_auth_model.AdminAuthDomainInterface) *rest_err.RestErr
	SignUp(admin admin_auth_model.AdminAuthDomainInterface) *rest_err.RestErr
}
