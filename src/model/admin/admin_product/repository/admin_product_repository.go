package admin_product_repository

import (
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	admin_product_model "matheuswww/coffeeShop-golang/src/model/admin/admin_product"
)

func NewAdminProductRepository() AdminProductRepository {
	return &adminProductRepository{}
}

type adminProductRepository struct{}

type AdminProductRepository interface {
	InsertProduct(AdminProductDomain admin_product_model.AdminProductDomainInterface) *rest_err.RestErr
}
