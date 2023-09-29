package product_repository

import (
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	product_model "matheuswww/coffeeShop-golang/src/model/product"
)

func NewProductDomainRepository() ProductRepository {
	return &productRepository{}
}

type productRepository struct {}

type ProductRepository interface {
	GetAll(*[]product_model.ProductDomainInterface) *rest_err.RestErr
}

