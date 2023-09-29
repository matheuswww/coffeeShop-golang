package admin_product_service

import (
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	admin_product_model "matheuswww/coffeeShop-golang/src/model/admin/admin_product"
	admin_product_repository "matheuswww/coffeeShop-golang/src/model/admin/admin_product/repository"
)

func NewAdminProductService(userRepository admin_product_repository.AdminProductRepository) AdminProductDomainService {
	return &adminProductDomainService{
		userRepository,
	}
}

type adminProductDomainService struct {
	repository admin_product_repository.AdminProductRepository
}

type AdminProductDomainService interface {
	InsertProduct(AdminProductDomain admin_product_model.AdminProductDomainInterface) *rest_err.RestErr
}
