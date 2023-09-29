package product_repository

import (
	"database/sql"
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	product_model "matheuswww/coffeeShop-golang/src/model/product"
)

func NewProductDomainRepository(database *sql.DB) ProductRepository {
	return &productRepository{
		database,
	}
}

type productRepository struct {
	database *sql.DB
}

type ProductRepository interface {
	GetAll(*[]product_model.ProductDomainInterface) *rest_err.RestErr
}
