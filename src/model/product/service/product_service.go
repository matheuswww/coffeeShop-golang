package product_service

import (
	"matheuswww/coffeeShop-golang/src/configuration/rest_err"
	product_model "matheuswww/coffeeShop-golang/src/model/product"
	product_repository "matheuswww/coffeeShop-golang/src/model/product/repository"
)

func NewProductService(repository product_repository.ProductRepository) ProductDomainRepository {
	return &productDomainService{
		repository: repository,
	}
}

type productDomainService struct {
	repository product_repository.ProductRepository
}

type ProductDomainRepository interface {
	GetAll(*[]product_model.ProductDomainInterface) *rest_err.RestErr
}