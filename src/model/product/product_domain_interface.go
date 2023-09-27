package product_model

type ProductDomainInterface interface{}

func NewProductDomain(name string, price float32, image []byte) ProductDomainInterface {
	return &productDomain{
		name,
		price,
		image,
	}
}
