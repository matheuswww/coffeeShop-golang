package admin_product_repository

import (
	"database/sql"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	admin_product_model "matheuswww/coffeeShop-golang/src/model/admin/admin_product"
)

func NewAdminProductRepository(database *sql.DB) AdminProductRepository {
	return &adminProductRepository{
		database,
	}
}

type adminProductRepository struct {
	database *sql.DB
}

type AdminProductRepository interface {
	InsertProduct(AdminProductDomain admin_product_model.AdminProductDomainInterface) *rest_err.RestErr
}
